package movie_service

import (
	"absolutecinema/internal/database/repository/movie"

	"github.com/google/uuid"
	"github.com/oapi-codegen/runtime/types"
)

type CreateMovieInput struct {
	Title       string      `json:"title"`
	Director    string      `json:"director"`
	Description string      `json:"description"`
	Photo       types.File  `json:"photo"`
	ActorIDs    []uuid.UUID `json:"actorIDs"`
}

type UpdateMovieInput struct {
	Title       string      `json:"title"`
	Director    string      `json:"director"`
	Description string      `json:"description"`
	Photo       types.File  `json:"photo"`
	ActorIDs    []uuid.UUID `json:"actorIDs"`
}

type MovieOutput struct {
	ID          uuid.UUID     `json:"id"`
	Title       string        `json:"title"`
	Director    string        `json:"director"`
	Description string        `json:"description"`
	PhotoPath   string        `json:"photoPath"`
	Actors      []ActorOutput `json:"actors"`
}

type ActorOutput struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func toMovieOutput(m *movie.Movie) *MovieOutput {
	actors := make([]ActorOutput, len(m.Actors))
	for i, a := range m.Actors {
		actors[i] = ActorOutput{
			ID:   a.ID,
			Name: a.Name,
		}
	}

	return &MovieOutput{
		ID:          m.ID,
		Title:       m.Title,
		Director:    m.Director,
		Description: m.Description,
		PhotoPath:   m.PhotoPath,
		Actors:      actors,
	}
}
