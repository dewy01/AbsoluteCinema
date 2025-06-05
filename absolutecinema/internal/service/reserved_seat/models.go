package reserved_seat_service

import "github.com/google/uuid"

type CreateReservedSeatInput struct {
	ReservationID uuid.UUID
	SeatID        uuid.UUID
}

type ReservedSeatOutput struct {
	ID            uuid.UUID
	ReservationID uuid.UUID
	SeatID        uuid.UUID
}
