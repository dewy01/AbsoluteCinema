package reservation

import (
	"absolutecinema/internal/database/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(res *Reservation) error
	GetByID(id uuid.UUID) (*Reservation, error)
	GetByUserID(userID uuid.UUID) ([]Reservation, error)
	UpdatePDF(id uuid.UUID, path string) error
	Delete(id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(res *Reservation) error {
	dbRes := ToDBReservation(res)
	if dbRes.ID == uuid.Nil {
		dbRes.ID = uuid.New()
	}
	for i := range dbRes.ReservedSeats {
		if dbRes.ReservedSeats[i].ID == uuid.Nil {
			dbRes.ReservedSeats[i].ID = uuid.New()
		}
	}
	return r.db.Create(dbRes).Error
}

func (r *repository) GetByID(id uuid.UUID) (*Reservation, error) {
	var dbRes models.Reservation
	if err := r.db.Preload("ReservedSeats").First(&dbRes, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return ToDomainReservation(&dbRes), nil
}

func (r *repository) GetByUserID(userID uuid.UUID) ([]Reservation, error) {
	var dbRes []models.Reservation
	if err := r.db.Preload("ReservedSeats").Where("user_id = ?", userID).Find(&dbRes).Error; err != nil {
		return nil, err
	}
	res := make([]Reservation, len(dbRes))
	for i := range dbRes {
		res[i] = *ToDomainReservation(&dbRes[i])
	}
	return res, nil
}

func (r *repository) UpdatePDF(id uuid.UUID, path string) error {
	return r.db.Model(&models.Reservation{}).Where("id = ?", id).Update("pdf_path", path).Error
}

func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Reservation{}, "id = ?", id).Error
}
