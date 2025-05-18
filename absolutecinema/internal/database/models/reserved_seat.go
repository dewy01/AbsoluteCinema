package models

import "github.com/google/uuid"

type ReservedSeat struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ReservationID uuid.UUID
	SeatID        uuid.UUID
}
