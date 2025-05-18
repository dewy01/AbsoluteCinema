package reservedseat

import (
	"absolutecinema/internal/database/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(rs *ReservedSeat) error
	GetByID(id uuid.UUID) (*ReservedSeat, error)
	GetByReservationID(reservationID uuid.UUID) ([]ReservedSeat, error)
	Delete(id uuid.UUID) error
	DeleteByReservationID(reservationID uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewReservedSeatRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(rs *ReservedSeat) error {
	dbRS := ToDBReservedSeat(rs)
	if dbRS.ID == uuid.Nil {
		dbRS.ID = uuid.New()
	}
	return r.db.Create(dbRS).Error
}

func (r *repository) GetByID(id uuid.UUID) (*ReservedSeat, error) {
	var dbRS models.ReservedSeat
	if err := r.db.First(&dbRS, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return ToDomainReservedSeat(&dbRS), nil
}

func (r *repository) GetByReservationID(reservationID uuid.UUID) ([]ReservedSeat, error) {
	var dbSeats []models.ReservedSeat
	if err := r.db.Where("reservation_id = ?", reservationID).Find(&dbSeats).Error; err != nil {
		return nil, err
	}

	seats := make([]ReservedSeat, len(dbSeats))
	for i, dbRS := range dbSeats {
		seats[i] = *ToDomainReservedSeat(&dbRS)
	}
	return seats, nil
}

func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.ReservedSeat{}, "id = ?", id).Error
}

func (r *repository) DeleteByReservationID(reservationID uuid.UUID) error {
	return r.db.Where("reservation_id = ?", reservationID).Delete(&models.ReservedSeat{}).Error
}
