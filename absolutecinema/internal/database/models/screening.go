package models

import (
	"time"

	"github.com/google/uuid"
)

type Screening struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	MovieID      uuid.UUID `gorm:"type:uuid"`
	Movie        Movie     `gorm:"foreignKey:MovieID"`
	RoomID       uuid.UUID `gorm:"type:uuid"`
	Room         Room      `gorm:"foreignKey:RoomID"`
	StartTime    time.Time
	Reservations []Reservation `gorm:"foreignKey:ScreeningID"`
}
