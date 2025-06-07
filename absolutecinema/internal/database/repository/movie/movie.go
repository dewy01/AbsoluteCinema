package movie

import (
	"absolutecinema/internal/database/models"

	"github.com/google/uuid"
)

type Movie struct {
	ID          uuid.UUID
	Title       string
	Director    string
	Description string
	PhotoPath   string
	Actors      []Actor
}

type Actor struct {
	ID   uuid.UUID
	Name string
}

type CreateMovie struct {
	ID          uuid.UUID
	Title       string
	Director    string
	Description string
	PhotoPath   string
	ActorIDs    []uuid.UUID
}

func ToDBCreateMovie(m *CreateMovie) *models.Movie {
	return &models.Movie{
		ID:          m.ID,
		Title:       m.Title,
		Director:    m.Director,
		Description: m.Description,
		PhotoPath:   m.PhotoPath,
	}
}

func ToDBMovie(m *Movie) *models.Movie {
	return &models.Movie{
		ID:          m.ID,
		Title:       m.Title,
		Director:    m.Director,
		Description: m.Description,
		PhotoPath:   m.PhotoPath,
	}
}

func ToDomainMovie(m *models.Movie) *Movie {
	actors := make([]Actor, len(m.Actors))
	for i, a := range m.Actors {
		actors[i] = Actor{ID: a.ID, Name: a.Name}
	}
	return &Movie{
		ID:          m.ID,
		Title:       m.Title,
		Director:    m.Director,
		Description: m.Description,
		PhotoPath:   m.PhotoPath,
		Actors:      actors,
	}
}
