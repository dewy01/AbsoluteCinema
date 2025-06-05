package reservation_service

import (
	reservedseat "absolutecinema/internal/database/repository/reserved_seat"

	"github.com/google/uuid"
)

type CreateReservationInput struct {
	ScreeningID   uuid.UUID
	UserID        *uuid.UUID
	GuestName     string
	GuestEmail    string
	ReservedSeats []reservedseat.ReservedSeat
}

type ReservedSeatOutput struct {
	ID     uuid.UUID
	SeatID uuid.UUID
}

type ReservationOutput struct {
	ID            uuid.UUID
	ScreeningID   uuid.UUID
	UserID        *uuid.UUID
	GuestName     string
	GuestEmail    string
	PDFPath       string
	ReservedSeats []ReservedSeatOutput
}
