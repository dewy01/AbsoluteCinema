package reservation_service

import (
	reservedseat "absolutecinema/internal/database/repository/reserved_seat"

	"github.com/google/uuid"
)

type CreateReservationInput struct {
	ScreeningID   uuid.UUID                   `json:"screeningID"`
	UserID        *uuid.UUID                  `json:"userID,omitempty"`
	GuestName     string                      `json:"guestName"`
	GuestEmail    string                      `json:"guestEmail,omitempty"`
	ReservedSeats []reservedseat.ReservedSeat `json:"reservedSeats"`
}

type UpdateReservationInput struct {
	ID            uuid.UUID
	UserID        *uuid.UUID
	GuestName     string
	GuestEmail    string
	ReservedSeats []reservedseat.ReservedSeat
}

type ReservedSeat = reservedseat.ReservedSeat

type ReservedSeatOutput struct {
	ID     uuid.UUID `json:"id"`
	SeatID uuid.UUID `json:"seatID"`
}

type ReservationOutput struct {
	ID            uuid.UUID            `json:"id"`
	ScreeningID   uuid.UUID            `json:"screeningID"`
	UserID        *uuid.UUID           `json:"userID,omitempty"`
	GuestName     string               `json:"guestName"`
	GuestEmail    string               `json:"guestEmail,omitempty"`
	PDFPath       string               `json:"pdfPath,omitempty"`
	ReservedSeats []ReservedSeatOutput `json:"reservedSeats,omitempty"`
}
