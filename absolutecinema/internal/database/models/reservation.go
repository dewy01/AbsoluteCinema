package models

import "github.com/google/uuid"

type Reservation struct {
	ID            uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ScreeningID   uuid.UUID  `gorm:"type:uuid"`
	Screening     Screening  `gorm:"foreignKey:ScreeningID"`
	UserID        *uuid.UUID `gorm:"type:uuid"`
	User          *User      `gorm:"foreignKey:UserID"`
	GuestName     string
	GuestEmail    string
	PDFPath       string
	ReservedSeats []ReservedSeat `gorm:"foreignKey:ReservationID"`
}
