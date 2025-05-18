package seat

import (
	"absolutecinema/internal/database/models"

	"github.com/google/uuid"
)

type Seat struct {
	ID     uuid.UUID
	Row    string
	Number int
	RoomID uuid.UUID
}

func ToDBSeat(s *Seat) *models.Seat {
	return &models.Seat{
		ID:     s.ID,
		Row:    s.Row,
		Number: s.Number,
		RoomID: s.RoomID,
	}
}

func ToDomainSeat(s *models.Seat) *Seat {
	return &Seat{
		ID:     s.ID,
		Row:    s.Row,
		Number: s.Number,
		RoomID: s.RoomID,
	}
}
