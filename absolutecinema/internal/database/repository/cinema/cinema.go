package cinema

import (
	"absolutecinema/internal/database/models"

	"github.com/google/uuid"
)

type Cinema struct {
	ID      uuid.UUID
	Name    string
	Address string
	RoomIDs []uuid.UUID
}

func ToDBCinema(c *Cinema) *models.Cinema {
	return &models.Cinema{
		ID:      c.ID,
		Name:    c.Name,
		Address: c.Address,
	}
}

func ToDomainCinema(c *models.Cinema) *Cinema {
	roomIDs := make([]uuid.UUID, len(c.Rooms))
	for i, r := range c.Rooms {
		roomIDs[i] = r.ID
	}
	return &Cinema{
		ID:      c.ID,
		Name:    c.Name,
		Address: c.Address,
		RoomIDs: roomIDs,
	}
}
