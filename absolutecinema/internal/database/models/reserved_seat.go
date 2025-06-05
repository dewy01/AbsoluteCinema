package models

import "github.com/google/uuid"

type ReservedSeat struct {
	ID            uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ReservationID uuid.UUID   `gorm:"type:uuid"`
	Reservation   Reservation `gorm:"foreignKey:ReservationID"`
	SeatID        uuid.UUID   `gorm:"type:uuid"`
	Seat          Seat        `gorm:"foreignKey:SeatID"`
}
