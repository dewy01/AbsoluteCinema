package reservation_service

import (
	"absolutecinema/internal/database/repository/reservation"
	reservedseat "absolutecinema/internal/database/repository/reserved_seat"
	"absolutecinema/internal/database/repository/screening"
	"absolutecinema/internal/database/repository/seat"
	"absolutecinema/pkg/fsystem"
	"bytes"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/google/uuid"
)

type Service interface {
	Create(input CreateReservationInput) (*ReservationOutput, error)
	Update(input UpdateReservationInput) (*ReservationOutput, error)
	GetByID(id uuid.UUID) (*ReservationOutput, error)
	GetByUserID(userID uuid.UUID) ([]ReservationOutput, error)
	UpdatePDF(id uuid.UUID, pdfPath string) error
	Delete(id uuid.UUID) error
}

type service struct {
	repo      reservation.Repository
	screening screening.Repository
	seat      seat.Repository
	fs        fsystem.FileStorage
}

func NewReservationService(repo reservation.Repository, screening screening.Repository, seat seat.Repository, fs fsystem.FileStorage) *service {
	return &service{
		repo:      repo,
		screening: screening,
		seat:      seat,
		fs:        fs,
	}
}

func (s *service) Create(input CreateReservationInput) (*ReservationOutput, error) {
	if input.ScreeningID == uuid.Nil {
		return nil, errors.New("screening ID is required")
	}
	if input.UserID == nil && (input.GuestName == "" || input.GuestEmail == "") {
		return nil, errors.New("guest name and email are required for guest reservations")
	}

	id := uuid.New()
	res := &reservation.Reservation{
		ID:            id,
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

	pdfBytes, err := s.generateReservationPDF(res.ID, input.GuestName, input.GuestEmail, input.ReservedSeats)
	if err != nil {
		return nil, fmt.Errorf("generate pdf: %w", err)
	}

	storagePath := id.String()
	filename := "reservation-" + id.String() + ".pdf"

	err = s.fs.SaveReservationFile(storagePath, filename, bytes.NewReader(pdfBytes))
	if err != nil {
		return nil, fmt.Errorf("save pdf: %w", err)
	}

	fullPath := filepath.Join("reservations", storagePath, filename)

	if err := s.repo.UpdatePDF(id, fullPath); err != nil {
		return nil, fmt.Errorf("update pdf path in db: %w", err)
	}

	updatedRes, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return toOutput(updatedRes), nil
}

func (s *service) Update(input UpdateReservationInput) (*ReservationOutput, error) {
	if input.UserID == nil && (input.GuestName == "" || input.GuestEmail == "") {
		return nil, errors.New("guest name and email are required for guest reservations")
	}

	before, err := s.repo.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	res := &reservation.Reservation{
		ID:            input.ID,
		ScreeningID:   before.ScreeningID,
		UserID:        input.UserID,
		GuestName:     input.GuestName,
		GuestEmail:    input.GuestEmail,
		PDFPath:       before.PDFPath,
		ReservedSeats: input.ReservedSeats,
	}

	if err := s.repo.Update(res); err != nil {
		return nil, fmt.Errorf("update reservation: %w", err)
	}

	pdfBytes, err := s.generateReservationPDF(input.ID, input.GuestName, input.GuestEmail, input.ReservedSeats)
	if err != nil {
		return nil, fmt.Errorf("generate pdf: %w", err)
	}

	storagePath := input.ID.String()
	filename := "reservation-" + input.ID.String() + ".pdf"

	err = s.fs.SaveReservationFile(storagePath, filename, bytes.NewReader(pdfBytes))
	if err != nil {
		return nil, fmt.Errorf("save pdf: %w", err)
	}

	fullPath := filepath.Join("reservations", storagePath, filename)

	if err := s.repo.UpdatePDF(input.ID, fullPath); err != nil {
		return nil, fmt.Errorf("update pdf path in db: %w", err)
	}

	updatedRes, err := s.repo.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	return toOutput(updatedRes), nil
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

func (s *service) generateReservationPDF(reservationID uuid.UUID, guestName, guestEmail string, reservedSeats []reservedseat.ReservedSeat) ([]byte, error) {
	seats := []SeatFile{}
	for _, rseat := range reservedSeats {
		seat, err := s.seat.GetByID(rseat.SeatID)
		if err != nil {
			return nil, err
		}
		seats = append(seats, SeatFile{
			Number: seat.Number,
			Row:    seat.Row,
		})
	}

	reservation, err := s.repo.GetByID(reservationID)
	if err != nil {
		return nil, err
	}

	screening, err := s.screening.GetByID(reservation.ScreeningID)
	if err != nil {
		return nil, err
	}

	resFile := &ReservationFile{
		ID:         reservationID,
		GuestName:  guestName,
		GuestEmail: guestEmail,
		Movie: MovieFile{
			Title:       screening.Movie.Title,
			Description: screening.Movie.Description,
		},
		StartTime:     screening.StartTime,
		ReservedSeats: seats,
		Room:          screening.Room.Name,
	}

	return generateReservationPDF(resFile)
}
