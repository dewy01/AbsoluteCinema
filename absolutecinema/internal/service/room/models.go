package room_service

import (
	"absolutecinema/internal/database/repository/room"

	"github.com/google/uuid"
)

type CreateRoomInput struct {
	Name     string
	CinemaID uuid.UUID
	Seats    []SeatInput
}

type UpdateRoomInput struct {
	ID       uuid.UUID
	Name     string
	CinemaID uuid.UUID
}

type RoomOutput struct {
	ID       uuid.UUID
	Name     string
	CinemaID uuid.UUID
	Seats    []SeatOutput
}

type SeatInput struct {
	Row    string
	Number int
}

type SeatOutput struct {
	ID     uuid.UUID
	Row    string
	Number int
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
