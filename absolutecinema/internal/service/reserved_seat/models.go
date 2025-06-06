package reserved_seat_service

import "github.com/google/uuid"

type CreateReservedSeatInput struct {
	ReservationID uuid.UUID `json:"reservationID"`
	SeatID        uuid.UUID `json:"seatID"`
}

type ReservedSeatOutput struct {
	ID            uuid.UUID `json:"id"`
	ReservationID uuid.UUID `json:"reservationID"`
	SeatID        uuid.UUID `json:"seatID"`
}
