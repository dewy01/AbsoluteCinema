package screening

import (
	"absolutecinema/internal/database/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(s *Screening) error
	GetByID(id uuid.UUID) (*Screening, error)
	GetByMovie(movieID uuid.UUID) ([]Screening, error)
	GetByRoom(roomID uuid.UUID) ([]Screening, error)
	GetAll() ([]Screening, error)
	Delete(id uuid.UUID) error
	UpdateStartTime(id uuid.UUID, newTime time.Time) error
}

type repository struct {
	db *gorm.DB
}

func NewScreeningRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(s *Screening) error {
	dbS := ToDBScreening(s)
	if dbS.ID == uuid.Nil {
		dbS.ID = uuid.New()
	}
	return r.db.Create(dbS).Error
}

func (r *repository) GetByID(id uuid.UUID) (*Screening, error) {
	var dbS models.Screening
	if err := r.db.First(&dbS, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return ToDomainScreening(&dbS), nil
}

func (r *repository) GetByMovie(movieID uuid.UUID) ([]Screening, error) {
	var screenings []models.Screening
	if err := r.db.Where("movie_id = ?", movieID).Find(&screenings).Error; err != nil {
		return nil, err
	}
	return mapScreenings(screenings), nil
}

func (r *repository) GetByRoom(roomID uuid.UUID) ([]Screening, error) {
	var screenings []models.Screening
	if err := r.db.Where("room_id = ?", roomID).Find(&screenings).Error; err != nil {
		return nil, err
	}
	return mapScreenings(screenings), nil
}

func (r *repository) GetAll() ([]Screening, error) {
	var screenings []models.Screening
	if err := r.db.Find(&screenings).Error; err != nil {
		return nil, err
	}
	return mapScreenings(screenings), nil
}

func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Screening{}, "id = ?", id).Error
}

func (r *repository) UpdateStartTime(id uuid.UUID, newTime time.Time) error {
	return r.db.Model(&models.Screening{}).Where("id = ?", id).Update("start_time", newTime).Error
}
