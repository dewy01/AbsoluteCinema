package models

import (
	"time"

	"github.com/google/uuid"
)

type Screening struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	MovieID      uuid.UUID
	RoomID       uuid.UUID
	StartTime    time.Time
	Reservations []Reservation
}
