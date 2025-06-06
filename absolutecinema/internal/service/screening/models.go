package screening_service

import (
	"absolutecinema/internal/database/repository/screening"
	"time"

	"github.com/google/uuid"
)

type MovieOutput struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Director    string    `json:"director"`
	Description string    `json:"description"`
	PhotoPath   string    `json:"photoPath"`
}

type RoomOutput struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type CreateScreeningInput struct {
	MovieID   uuid.UUID `json:"movieID"`
	RoomID    uuid.UUID `json:"roomID"`
	StartTime time.Time `json:"startTime"`
}

type ScreeningOutput struct {
	ID        uuid.UUID   `json:"id"`
	StartTime time.Time   `json:"startTime"`
	Movie     MovieOutput `json:"movie"`
	Room      RoomOutput  `json:"room"`
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
