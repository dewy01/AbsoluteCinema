package reservation_service

import (
	"absolutecinema/internal/database/repository/reservation"
	"errors"

	"github.com/google/uuid"
)

type Service interface {
	Create(input CreateReservationInput) (*ReservationOutput, error)
	GetByID(id uuid.UUID) (*ReservationOutput, error)
	GetByUserID(userID uuid.UUID) ([]ReservationOutput, error)
	UpdatePDF(id uuid.UUID, pdfPath string) error
	Delete(id uuid.UUID) error
}

type service struct {
	repo reservation.Repository
}

func NewReservationService(repo reservation.Repository) *service {
	return &service{repo: repo}
}

func (s *service) Create(input CreateReservationInput) (*ReservationOutput, error) {
	if input.ScreeningID == uuid.Nil {
		return nil, errors.New("screening ID is required")
	}
	if input.UserID == nil && (input.GuestName == "" || input.GuestEmail == "") {
		return nil, errors.New("guest name and email are required for guest reservations")
	}

	res := &reservation.Reservation{
		ID:            uuid.New(),
		ScreeningID:   input.ScreeningID,
		UserID:        input.UserID,
		GuestName:     input.GuestName,
		GuestEmail:    input.GuestEmail,
		PDFPath:       "",
		ReservedSeats: input.ReservedSeats,
	}

	if err := s.repo.Create(res); err != nil {
		return nil, err
	}

	return toOutput(res), nil
}

func (s *service) GetByID(id uuid.UUID) (*ReservationOutput, error) {
	res, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return toOutput(res), nil
}

func (s *service) GetByUserID(userID uuid.UUID) ([]ReservationOutput, error) {
	reservations, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	outputs := make([]ReservationOutput, len(reservations))
	for i, r := range reservations {
		outputs[i] = *toOutput(&r)
	}
	return outputs, nil
}

func (s *service) UpdatePDF(id uuid.UUID, pdfPath string) error {
	if pdfPath == "" {
		return errors.New("pdf path cannot be empty")
	}
	return s.repo.UpdatePDF(id, pdfPath)
}

func (s *service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func toOutput(r *reservation.Reservation) *ReservationOutput {
	seats := make([]ReservedSeatOutput, len(r.ReservedSeats))
	for i, seat := range r.ReservedSeats {
		seats[i] = ReservedSeatOutput{
			ID:     seat.ID,
			SeatID: seat.SeatID,
		}
	}
	return &ReservationOutput{
		ID:            r.ID,
		ScreeningID:   r.ScreeningID,
		UserID:        r.UserID,
		GuestName:     r.GuestName,
		GuestEmail:    r.GuestEmail,
		PDFPath:       r.PDFPath,
		ReservedSeats: seats,
	}
}
