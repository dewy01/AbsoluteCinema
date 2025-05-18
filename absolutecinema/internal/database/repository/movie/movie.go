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
	ActorIDs    []uuid.UUID
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
	actorIDs := make([]uuid.UUID, len(m.Actors))
	for i, actor := range m.Actors {
		actorIDs[i] = actor.ID
	}
	return &Movie{
		ID:          m.ID,
		Title:       m.Title,
		Director:    m.Director,
		Description: m.Description,
		PhotoPath:   m.PhotoPath,
		ActorIDs:    actorIDs,
	}
}
