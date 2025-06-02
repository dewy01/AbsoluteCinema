package cinema_service

import (
	"absolutecinema/internal/database/repository/cinema"
	"errors"

	"github.com/google/uuid"
)

type Service interface {
	Create(input CreateCinemaInput) (*CinemaOutput, error)
	GetByID(id uuid.UUID) (*CinemaOutput, error)
	GetAll() ([]CinemaOutput, error)
	Update(id uuid.UUID, input UpdateCinemaInput) (*CinemaOutput, error)
	Delete(id uuid.UUID) error
}

type service struct {
	repo cinema.Repository
}

func NewCinemaService(repo cinema.Repository) *service {
	return &service{repo: repo}
}

func (s *service) Create(input CreateCinemaInput) (*CinemaOutput, error) {
	if input.Name == "" {
		return nil, errors.New("name is required")
	}
	if input.Address == "" {
		return nil, errors.New("address is required")
	}

	cinemaDomain := &cinema.Cinema{
		ID:      uuid.New(),
		Name:    input.Name,
		Address: input.Address,
		RoomIDs: input.RoomIDs,
	}

	err := s.repo.Create(cinemaDomain)
	if err != nil {
		return nil, err
	}

	return &CinemaOutput{
		ID:      cinemaDomain.ID,
		Name:    cinemaDomain.Name,
		Address: cinemaDomain.Address,
		RoomIDs: cinemaDomain.RoomIDs,
	}, nil
}

func (s *service) GetByID(id uuid.UUID) (*CinemaOutput, error) {
	cinemaDomain, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &CinemaOutput{
		ID:      cinemaDomain.ID,
		Name:    cinemaDomain.Name,
		Address: cinemaDomain.Address,
		RoomIDs: cinemaDomain.RoomIDs,
	}, nil
}

func (s *service) GetAll() ([]CinemaOutput, error) {
	cinemas, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	outputs := make([]CinemaOutput, len(cinemas))
	for i, c := range cinemas {
		outputs[i] = CinemaOutput{
			ID:      c.ID,
			Name:    c.Name,
			Address: c.Address,
			RoomIDs: c.RoomIDs,
		}
	}

	return outputs, nil
}

func (s *service) Update(id uuid.UUID, input UpdateCinemaInput) (*CinemaOutput, error) {
	cinemaDomain, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if input.Name != "" {
		cinemaDomain.Name = input.Name
	}
	if input.Address != "" {
		cinemaDomain.Address = input.Address
	}
	if input.RoomIDs != nil {
		cinemaDomain.RoomIDs = input.RoomIDs
	}

	if err := s.repo.Update(cinemaDomain); err != nil {
		return nil, err
	}

	return &CinemaOutput{
		ID:      cinemaDomain.ID,
		Name:    cinemaDomain.Name,
		Address: cinemaDomain.Address,
		RoomIDs: cinemaDomain.RoomIDs,
	}, nil
}

func (s *service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
