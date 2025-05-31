package cinema_service

import "github.com/google/uuid"

type CreateCinemaInput struct {
	Name    string
	Address string
	RoomIDs []uuid.UUID
}

type UpdateCinemaInput struct {
	Name    string
	Address string
	RoomIDs []uuid.UUID
}

type CinemaOutput struct {
	ID      uuid.UUID
	Name    string
	Address string
	RoomIDs []uuid.UUID
}
