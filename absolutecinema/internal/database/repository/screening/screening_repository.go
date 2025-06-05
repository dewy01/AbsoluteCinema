package screening

import (
	"absolutecinema/internal/database/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(s *ScreeningInput) error
	GetByID(id uuid.UUID) (*Screening, error)
	GetByMovie(movieID uuid.UUID, day *time.Time) ([]Screening, error)
	GetByRoom(roomID uuid.UUID, day *time.Time) ([]Screening, error)
	GetByCinema(cinemaID uuid.UUID, day *time.Time) ([]Screening, error)
	GetAll(day *time.Time) ([]Screening, error)
	Delete(id uuid.UUID) error
	UpdateStartTime(id uuid.UUID, newTime time.Time) error
}

type repository struct {
	db *gorm.DB
}

func NewScreeningRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(s *ScreeningInput) error {
	dbS := ToDBScreening(s)
	if dbS.ID == uuid.Nil {
		dbS.ID = uuid.New()
	}
	return r.db.Create(dbS).Error
}

func (r *repository) GetByID(id uuid.UUID) (*Screening, error) {
	var dbS models.Screening
	err := r.db.
		Preload("Movie").
		Preload("Room").
		Preload("Room.Cinema").
		Preload("Reservations.ReservedSeats").
		First(&dbS, "id = ?", id).Error

	if err != nil {
		return nil, err
	}
	return ToDomainScreening(&dbS), nil
}

func (r *repository) GetByMovie(movieID uuid.UUID, day *time.Time) ([]Screening, error) {
	var screenings []models.Screening
	query := r.db.
		Where("movie_id = ?", movieID).
		Preload("Movie").
		Preload("Room").
		Preload("Room.Cinema")

	if day != nil {
		start, end := dayBounds(*day)
		query = query.Where("start_time BETWEEN ? AND ?", start, end)
	}

	if err := query.Find(&screenings).Error; err != nil {
		return nil, err
	}
	return mapScreenings(screenings), nil
}

func (r *repository) GetByRoom(roomID uuid.UUID, day *time.Time) ([]Screening, error) {
	var screenings []models.Screening
	query := r.db.
		Where("room_id = ?", roomID).
		Preload("Movie").
		Preload("Room").
		Preload("Room.Cinema")

	if day != nil {
		start, end := dayBounds(*day)
		query = query.Where("start_time BETWEEN ? AND ?", start, end)
	}

	if err := query.Find(&screenings).Error; err != nil {
		return nil, err
	}
	return mapScreenings(screenings), nil
}

func (r *repository) GetByCinema(cinemaID uuid.UUID, day *time.Time) ([]Screening, error) {
	var screenings []models.Screening
	query := r.db.
		Joins("JOIN rooms ON rooms.id = screenings.room_id").
		Where("rooms.cinema_id = ?", cinemaID).
		Preload("Movie").
		Preload("Room").
		Preload("Room.Cinema")

	if day != nil {
		start, end := dayBounds(*day)
		query = query.Where("start_time BETWEEN ? AND ?", start, end)
	}

	if err := query.Find(&screenings).Error; err != nil {
		return nil, err
	}

	return mapScreenings(screenings), nil
}

func (r *repository) GetAll(day *time.Time) ([]Screening, error) {
	var screenings []models.Screening
	query := r.db.
		Preload("Movie").
		Preload("Room").
		Preload("Room.Cinema")

	if day != nil {
		start, end := dayBounds(*day)
		query = query.Where("start_time BETWEEN ? AND ?", start, end)
	}

	if err := query.Find(&screenings).Error; err != nil {
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
