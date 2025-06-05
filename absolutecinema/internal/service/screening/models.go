package screening_service

import (
	"absolutecinema/internal/database/repository/screening"
	"time"

	"github.com/google/uuid"
)

type CreateScreeningInput struct {
	MovieID   uuid.UUID
	RoomID    uuid.UUID
	StartTime time.Time
}

type ScreeningOutput struct {
	ID        uuid.UUID
	MovieID   uuid.UUID
	RoomID    uuid.UUID
	StartTime time.Time
}

func toOutput(s *screening.Screening) *ScreeningOutput {
	return &ScreeningOutput{
		ID:        s.ID,
		MovieID:   s.MovieID,
		RoomID:    s.RoomID,
		StartTime: s.StartTime,
	}
}
