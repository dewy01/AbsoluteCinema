package actor

import (
	"absolutecinema/internal/database/models"

	"github.com/google/uuid"
)

type Actor struct {
	ID       uuid.UUID
	Name     string
	MovieIDs []uuid.UUID
}

func ToDBActor(a *Actor) *models.Actor {
	return &models.Actor{
		ID:   a.ID,
		Name: a.Name,
	}
}

func ToDomainActor(a *models.Actor) *Actor {
	movieIDs := make([]uuid.UUID, len(a.Movies))
	for i, m := range a.Movies {
		movieIDs[i] = m.ID
	}
	return &Actor{
		ID:       a.ID,
		Name:     a.Name,
		MovieIDs: movieIDs,
	}
}
