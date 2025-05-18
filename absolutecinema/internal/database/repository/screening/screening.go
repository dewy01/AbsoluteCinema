package screening

import (
	"absolutecinema/internal/database/models"
	"time"

	"github.com/google/uuid"
)

type Screening struct {
	ID        uuid.UUID
	MovieID   uuid.UUID
	RoomID    uuid.UUID
	StartTime time.Time
}

func ToDBScreening(s *Screening) *models.Screening {
	return &models.Screening{
		ID:        s.ID,
		MovieID:   s.MovieID,
		RoomID:    s.RoomID,
		StartTime: s.StartTime,
	}
}

func ToDomainScreening(s *models.Screening) *Screening {
	return &Screening{
		ID:        s.ID,
		MovieID:   s.MovieID,
		RoomID:    s.RoomID,
		StartTime: s.StartTime,
	}
}

func mapScreenings(src []models.Screening) []Screening {
	result := make([]Screening, len(src))
	for i, s := range src {
		result[i] = *ToDomainScreening(&s)
	}
	return result
}
