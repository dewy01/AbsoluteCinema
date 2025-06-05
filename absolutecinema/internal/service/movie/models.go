package movie_service

import (
	"github.com/google/uuid"
	"github.com/oapi-codegen/runtime/types"
)

type CreateMovieInput struct {
	Title       string
	Director    string
	Description string
	Photo       types.File
	ActorIDs    []uuid.UUID
}

type UpdateMovieInput struct {
	Title       string
	Director    string
	Description string
	Photo       types.File
	ActorIDs    []uuid.UUID
}

type MovieOutput struct {
	ID          uuid.UUID
	Title       string
	Director    string
	Description string
	PhotoPath   string
	ActorIDs    []uuid.UUID
}
