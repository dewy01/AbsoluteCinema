package room_service

import (
	"absolutecinema/internal/database/repository/room"

	"github.com/google/uuid"
)

type CreateRoomInput struct {
	Name     string      `json:"name"`
	CinemaID uuid.UUID   `json:"cinemaID"`
	Seats    []SeatInput `json:"seats"`
}

type UpdateRoomInput struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	CinemaID uuid.UUID `json:"cinemaID"`
}

type RoomOutput struct {
	ID       uuid.UUID    `json:"id"`
	Name     string       `json:"name"`
	CinemaID uuid.UUID    `json:"cinemaID"`
	Seats    []SeatOutput `json:"seats"`
}

type SeatInput struct {
	Row    string `json:"row"`
	Number int    `json:"number"`
}

type SeatOutput struct {
	ID     uuid.UUID `json:"id"`
	Row    string    `json:"row"`
	Number int       `json:"number"`
}

func toRoomOutput(r *room.Room) *RoomOutput {
	seats := make([]SeatOutput, len(r.Seats))
	for i, s := range r.Seats {
		seats[i] = SeatOutput{
			ID:     s.ID,
			Row:    s.Row,
			Number: s.Number,
		}
	}

	return &RoomOutput{
		ID:       r.ID,
		Name:     r.Name,
		CinemaID: r.CinemaID,
		Seats:    seats,
	}
}
