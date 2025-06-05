package movie_service

import "github.com/google/uuid"

type CreateMovieInput struct {
	Title       string
	Director    string
	Description string
	PhotoPath   string
	ActorIDs    []uuid.UUID
}

type UpdateMovieInput struct {
	Title       string
	Director    string
	Description string
	PhotoPath   string
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
