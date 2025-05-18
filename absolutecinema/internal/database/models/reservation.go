package models

import "github.com/google/uuid"

type Reservation struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ScreeningID   uuid.UUID
	UserID        *uuid.UUID
	User          *User
	GuestName     string
	GuestEmail    string
	PDFPath       string
	ReservedSeats []ReservedSeat
}
