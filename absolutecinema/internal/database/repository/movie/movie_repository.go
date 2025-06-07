package movie

import (
	"absolutecinema/internal/database/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(m *CreateMovie) error
	Update(id uuid.UUID, m *CreateMovie) error
	GetByID(id uuid.UUID) (*Movie, error)
	GetAll() ([]Movie, error)
	Delete(id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(m *CreateMovie) error {
	dbM := ToDBCreateMovie(m)
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
		if err := r.db.Model(dbM).Association("Actors").Replace(&actors); err != nil {
			return err
		}
	}

	return nil
}

func (r *repository) Update(id uuid.UUID, m *CreateMovie) error {
	dbM := ToDBCreateMovie(m)

	if err := r.db.Model(&models.Movie{}).Where("id = ?", id).Updates(dbM).Error; err != nil {
		return err
	}

	var movie models.Movie
	if err := r.db.First(&movie, "id = ?", id).Error; err != nil {
		return err
	}

	if len(m.ActorIDs) > 0 {
		var actors []models.Actor
		if err := r.db.Where("id IN ?", m.ActorIDs).Find(&actors).Error; err != nil {
			return err
		}
		return r.db.Model(&movie).Association("Actors").Replace(&actors)
	} else {
		return r.db.Model(&movie).Association("Actors").Clear()
	}
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

func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Movie{}, "id = ?", id).Error
}
