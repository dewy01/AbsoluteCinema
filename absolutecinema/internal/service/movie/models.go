package movie_service

import (
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
	ID          uuid.UUID   `json:"id"`
	Title       string      `json:"title"`
	Director    string      `json:"director"`
	Description string      `json:"description"`
	PhotoPath   string      `json:"photoPath"`
	ActorIDs    []uuid.UUID `json:"actorIDs"`
}
