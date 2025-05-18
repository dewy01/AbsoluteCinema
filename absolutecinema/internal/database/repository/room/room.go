package room

import (
	"absolutecinema/internal/database/models"
	"absolutecinema/internal/database/repository/seat"

	"github.com/google/uuid"
)

type Room struct {
	ID       uuid.UUID
	Name     string
	CinemaID uuid.UUID
	Seats    []seat.Seat
}

func ToDBRoom(r *Room) *models.Room {
	seats := make([]models.Seat, len(r.Seats))
	for i, s := range r.Seats {
		seats[i] = *seat.ToDBSeat(&s)
	}
	return &models.Room{
		ID:       r.ID,
		Name:     r.Name,
		CinemaID: r.CinemaID,
		Seats:    seats,
	}
}

func ToDomainRoom(r *models.Room) *Room {
	seats := make([]seat.Seat, len(r.Seats))
	for i, s := range r.Seats {
		seats[i] = *seat.ToDomainSeat(&s)
	}
	return &Room{
		ID:       r.ID,
		Name:     r.Name,
		CinemaID: r.CinemaID,
		Seats:    seats,
	}
}
