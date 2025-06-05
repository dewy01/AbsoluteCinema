package room_service

import (
	"absolutecinema/internal/database/repository/room"
	"absolutecinema/internal/database/repository/seat"

	"errors"

	"github.com/google/uuid"
)

type Service interface {
	Create(input CreateRoomInput) (*RoomOutput, error)
	GetByID(id uuid.UUID) (*RoomOutput, error)
	GetByCinemaID(cinemaID uuid.UUID) ([]RoomOutput, error)
	Update(input UpdateRoomInput) error
	Delete(id uuid.UUID) error
}

type service struct {
	repo room.Repository
}

func NewRoomService(repo room.Repository) *service {
	return &service{repo: repo}
}

func (s *service) Create(input CreateRoomInput) (*RoomOutput, error) {
	if input.Name == "" || input.CinemaID == uuid.Nil {
		return nil, errors.New("name and cinema ID are required")
	}

	roomID := uuid.New()
	seats := make([]seat.Seat, len(input.Seats))
	for i, s := range input.Seats {
		seats[i] = seat.Seat{
			ID:     uuid.New(),
			Row:    s.Row,
			Number: s.Number,
		}
	}

	r := &room.Room{
		ID:       roomID,
		Name:     input.Name,
		CinemaID: input.CinemaID,
		Seats:    seats,
	}

	if err := s.repo.Create(r); err != nil {
		return nil, err
	}

	return toRoomOutput(r), nil
}

func (s *service) GetByID(id uuid.UUID) (*RoomOutput, error) {
	r, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return toRoomOutput(r), nil
}

func (s *service) GetByCinemaID(cinemaID uuid.UUID) ([]RoomOutput, error) {
	rooms, err := s.repo.GetByCinemaID(cinemaID)
	if err != nil {
		return nil, err
	}

	output := make([]RoomOutput, len(rooms))
	for i, r := range rooms {
		output[i] = *toRoomOutput(&r)
	}
	return output, nil
}

func (s *service) Update(input UpdateRoomInput) error {
	if input.ID == uuid.Nil || input.Name == "" || input.CinemaID == uuid.Nil {
		return errors.New("id, name and cinema ID are required")
	}
	return s.repo.Update(&room.Room{
		ID:       input.ID,
		Name:     input.Name,
		CinemaID: input.CinemaID,
	})
}

func (s *service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
