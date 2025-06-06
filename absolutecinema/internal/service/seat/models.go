package seat_service

import (
	"absolutecinema/internal/database/repository/seat"

	"github.com/google/uuid"
)

type CreateSeatInput struct {
	Row    string    `json:"row"`
	Number int       `json:"number"`
	RoomID uuid.UUID `json:"roomID"`
}

type UpdateSeatInput struct {
	ID     uuid.UUID `json:"id"`
	Row    string    `json:"row"`
	Number int       `json:"number"`
}

type SeatOutput struct {
	ID     uuid.UUID `json:"id"`
	Row    string    `json:"row"`
	Number int       `json:"number"`
	RoomID uuid.UUID `json:"roomID"`
}

type SeatWithReservationStatusOutput struct {
	ID         uuid.UUID `json:"id"`
	Row        string    `json:"row"`
	Number     int       `json:"number"`
	RoomID     uuid.UUID `json:"roomID"`
	IsReserved bool      `json:"isReserved"`
}

func toOutput(s *seat.Seat) *SeatOutput {
	return &SeatOutput{
		ID:     s.ID,
		Row:    s.Row,
		Number: s.Number,
		RoomID: s.RoomID,
	}
}
