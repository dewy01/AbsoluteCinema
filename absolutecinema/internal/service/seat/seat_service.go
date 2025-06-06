package seat_service

import (
	"absolutecinema/internal/database/repository/seat"
	"errors"

	"github.com/google/uuid"
)

type Service interface {
	Create(input CreateSeatInput) (*SeatOutput, error)
	GetByID(id uuid.UUID) (*SeatOutput, error)
	GetByRoomID(roomID uuid.UUID) ([]SeatOutput, error)
	GetByScreeningID(screeningID uuid.UUID) ([]SeatWithReservationStatusOutput, error)
	Update(input UpdateSeatInput) error
	Delete(id uuid.UUID) error
}

type service struct {
	repo seat.Repository
}

func NewSeatService(repo seat.Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(input CreateSeatInput) (*SeatOutput, error) {
	if input.RoomID == uuid.Nil || input.Row == "" || input.Number <= 0 {
		return nil, errors.New("invalid seat input")
	}

	newSeat := &seat.Seat{
		ID:     uuid.New(),
		Row:    input.Row,
		Number: input.Number,
		RoomID: input.RoomID,
	}

	if err := s.repo.Create(newSeat); err != nil {
		return nil, err
	}

	return toOutput(newSeat), nil
}

func (s *service) GetByID(id uuid.UUID) (*SeatOutput, error) {
	if id == uuid.Nil {
		return nil, errors.New("invalid seat ID")
	}

	result, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return toOutput(result), nil
}

func (s *service) GetByRoomID(roomID uuid.UUID) ([]SeatOutput, error) {
	if roomID == uuid.Nil {
		return nil, errors.New("invalid room ID")
	}

	seats, err := s.repo.GetByRoomID(roomID)
	if err != nil {
		return nil, err
	}

	output := make([]SeatOutput, len(seats))
	for i, s := range seats {
		output[i] = *toOutput(&s)
	}
	return output, nil
}

func (s *service) GetByScreeningID(screeningID uuid.UUID) ([]SeatWithReservationStatusOutput, error) {
	if screeningID == uuid.Nil {
		return nil, errors.New("invalid screening ID")
	}

	seats, err := s.repo.GetByScreeningID(screeningID)
	if err != nil {
		return nil, err
	}

	output := make([]SeatWithReservationStatusOutput, len(seats))
	for i, seat := range seats {
		output[i] = SeatWithReservationStatusOutput{
			ID:         seat.ID,
			Row:        seat.Row,
			Number:     seat.Number,
			RoomID:     seat.RoomID,
			IsReserved: seat.IsReserved,
		}
	}
	return output, nil
}

func (s *service) Update(input UpdateSeatInput) error {
	if input.ID == uuid.Nil || input.Row == "" || input.Number <= 0 {
		return errors.New("invalid seat update input")
	}

	return s.repo.Update(&seat.Seat{
		ID:     input.ID,
		Row:    input.Row,
		Number: input.Number,
	})
}

func (s *service) Delete(id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("invalid seat ID")
	}
	return s.repo.Delete(id)
}
