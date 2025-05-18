package seat

import (
	"absolutecinema/internal/database/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(seat *Seat) error
	GetByID(id uuid.UUID) (*Seat, error)
	GetByRoomID(roomID uuid.UUID) ([]Seat, error)
	Update(seat *Seat) error
	Delete(id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewSeatRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(seat *Seat) error {
	dbSeat := ToDBSeat(seat)
	if dbSeat.ID == uuid.Nil {
		dbSeat.ID = uuid.New()
	}
	return r.db.Create(dbSeat).Error
}

func (r *repository) GetByID(id uuid.UUID) (*Seat, error) {
	var dbSeat models.Seat
	if err := r.db.First(&dbSeat, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return ToDomainSeat(&dbSeat), nil
}

func (r *repository) GetByRoomID(roomID uuid.UUID) ([]Seat, error) {
	var dbSeats []models.Seat
	if err := r.db.Where("room_id = ?", roomID).Find(&dbSeats).Error; err != nil {
		return nil, err
	}

	seats := make([]Seat, len(dbSeats))
	for i, s := range dbSeats {
		seats[i] = *ToDomainSeat(&s)
	}
	return seats, nil
}

func (r *repository) Update(seat *Seat) error {
	return r.db.Model(&models.Seat{}).Where("id = ?", seat.ID).Updates(map[string]interface{}{
		"row":    seat.Row,
		"number": seat.Number,
	}).Error
}

func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Seat{}, "id = ?", id).Error
}
