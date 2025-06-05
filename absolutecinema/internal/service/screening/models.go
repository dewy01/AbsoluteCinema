package screening_service

import (
	"absolutecinema/internal/database/repository/screening"
	"time"

	"github.com/google/uuid"
)

type MovieOutput struct {
	ID          uuid.UUID
	Title       string
	Director    string
	Description string
	PhotoPath   string
}

type RoomOutput struct {
	ID   uuid.UUID
	Name string
}

type CreateScreeningInput struct {
	MovieID   uuid.UUID
	RoomID    uuid.UUID
	StartTime time.Time
}

type ScreeningOutput struct {
	ID        uuid.UUID
	StartTime time.Time
	Movie     MovieOutput
	Room      RoomOutput
}

func toOutput(s *screening.Screening) *ScreeningOutput {
	return &ScreeningOutput{
		ID:        s.ID,
		StartTime: s.StartTime,
		Movie: MovieOutput{
			ID:          s.Movie.ID,
			Title:       s.Movie.Title,
			Director:    s.Movie.Director,
			Description: s.Movie.Description,
			PhotoPath:   s.Movie.PhotoPath,
		},
		Room: RoomOutput{
			ID:   s.Room.ID,
			Name: s.Room.Name,
		},
	}
}

func mapToOutput(src []screening.Screening) []ScreeningOutput {
	result := make([]ScreeningOutput, len(src))
	for i, s := range src {
		result[i] = *toOutput(&s)
	}
	return result
}
