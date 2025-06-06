package cinema_service

import "github.com/google/uuid"

type CreateCinemaInput struct {
	Name    string      `json:"name"`
	Address string      `json:"address"`
	RoomIDs []uuid.UUID `json:"roomIDs"`
}

type UpdateCinemaInput struct {
	Name    string      `json:"name"`
	Address string      `json:"address"`
	RoomIDs []uuid.UUID `json:"roomIDs"`
}

type CinemaOutput struct {
	ID      uuid.UUID   `json:"id"`
	Name    string      `json:"name"`
	Address string      `json:"address"`
	RoomIDs []uuid.UUID `json:"roomIDs"`
}
