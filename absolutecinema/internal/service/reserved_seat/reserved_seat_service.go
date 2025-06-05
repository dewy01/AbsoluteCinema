package reserved_seat_service

import (
	reservedseat "absolutecinema/internal/database/repository/reserved_seat"
	"errors"

	"github.com/google/uuid"
)

type Service interface {
	Create(input CreateReservedSeatInput) (*ReservedSeatOutput, error)
	GetByID(id uuid.UUID) (*ReservedSeatOutput, error)
	GetByReservationID(reservationID uuid.UUID) ([]ReservedSeatOutput, error)
	Delete(id uuid.UUID) error
	DeleteByReservationID(reservationID uuid.UUID) error
}

type service struct {
	repo reservedseat.Repository
}

func NewReservedSeatService(repo reservedseat.Repository) *service {
	return &service{repo: repo}
}

func (s *service) Create(input CreateReservedSeatInput) (*ReservedSeatOutput, error) {
	if input.SeatID == uuid.Nil || input.ReservationID == uuid.Nil {
		return nil, errors.New("seat ID and reservation ID are required")
	}

	rs := &reservedseat.ReservedSeat{
		ID:            uuid.New(),
		ReservationID: input.ReservationID,
		SeatID:        input.SeatID,
	}

	if err := s.repo.Create(rs); err != nil {
		return nil, err
	}

	return &ReservedSeatOutput{
		ID:            rs.ID,
		ReservationID: rs.ReservationID,
		SeatID:        rs.SeatID,
	}, nil
}

func (s *service) GetByID(id uuid.UUID) (*ReservedSeatOutput, error) {
	rs, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &ReservedSeatOutput{
		ID:            rs.ID,
		ReservationID: rs.ReservationID,
		SeatID:        rs.SeatID,
	}, nil
}

func (s *service) GetByReservationID(reservationID uuid.UUID) ([]ReservedSeatOutput, error) {
	seats, err := s.repo.GetByReservationID(reservationID)
	if err != nil {
		return nil, err
	}

	output := make([]ReservedSeatOutput, len(seats))
	for i, s := range seats {
		output[i] = ReservedSeatOutput{
			ID:            s.ID,
			ReservationID: s.ReservationID,
			SeatID:        s.SeatID,
		}
	}
	return output, nil
}

func (s *service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *service) DeleteByReservationID(reservationID uuid.UUID) error {
	return s.repo.DeleteByReservationID(reservationID)
}
