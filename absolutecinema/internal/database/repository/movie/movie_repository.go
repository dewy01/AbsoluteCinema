package movie

import (
	"absolutecinema/internal/database/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(m *Movie) error
	GetByID(id uuid.UUID) (*Movie, error)
	GetAll() ([]Movie, error)
	Update(m *Movie) error
	Delete(id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(m *Movie) error {
	dbM := ToDBMovie(m)
	if dbM.ID == uuid.Nil {
		dbM.ID = uuid.New()
	}
	if err := r.db.Create(dbM).Error; err != nil {
		return err
	}

	if len(m.ActorIDs) > 0 {
		var actors []models.Actor
		if err := r.db.Where("id IN ?", m.ActorIDs).Find(&actors).Error; err != nil {
			return err
		}
		return r.db.Model(dbM).Association("Actors").Replace(&actors)
	}

	return nil
}

func (r *repository) GetByID(id uuid.UUID) (*Movie, error) {
	var dbM models.Movie
	if err := r.db.Preload("Actors").First(&dbM, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return ToDomainMovie(&dbM), nil
}

func (r *repository) GetAll() ([]Movie, error) {
	var dbMovies []models.Movie
	if err := r.db.Preload("Actors").Find(&dbMovies).Error; err != nil {
		return nil, err
	}
	movies := make([]Movie, len(dbMovies))
	for i, m := range dbMovies {
		movies[i] = *ToDomainMovie(&m)
	}
	return movies, nil
}

func (r *repository) Update(m *Movie) error {
	dbM := ToDBMovie(m)
	if err := r.db.Model(&models.Movie{}).Where("id = ?", m.ID).Updates(dbM).Error; err != nil {
		return err
	}

	if len(m.ActorIDs) > 0 {
		var actors []models.Actor
		if err := r.db.Where("id IN ?", m.ActorIDs).Find(&actors).Error; err != nil {
			return err
		}
		return r.db.Model(dbM).Association("Actors").Replace(&actors)
	} else {
		return r.db.Model(dbM).Association("Actors").Clear()
	}
}

func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Movie{}, "id = ?", id).Error
}
