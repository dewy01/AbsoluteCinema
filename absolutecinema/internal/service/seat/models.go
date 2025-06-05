package seat_service

import (
	"absolutecinema/internal/database/repository/seat"

	"github.com/google/uuid"
)

type CreateSeatInput struct {
	Row    string
	Number int
	RoomID uuid.UUID
}

type UpdateSeatInput struct {
	ID     uuid.UUID
	Row    string
	Number int
}

type SeatOutput struct {
	ID     uuid.UUID
	Row    string
	Number int
	RoomID uuid.UUID
}

func toOutput(s *seat.Seat) *SeatOutput {
	return &SeatOutput{
		ID:     s.ID,
		Row:    s.Row,
		Number: s.Number,
		RoomID: s.RoomID,
	}
}
