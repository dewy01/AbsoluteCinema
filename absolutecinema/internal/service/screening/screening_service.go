package screening_service

import (
	"absolutecinema/internal/database/repository/screening"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	Create(input CreateScreeningInput) (*ScreeningOutput, error)
	GetByID(id uuid.UUID) (*ScreeningOutput, error)
	GetByMovie(movieID uuid.UUID, day *time.Time) ([]ScreeningOutput, error)
	GetByRoom(roomID uuid.UUID, day *time.Time) ([]ScreeningOutput, error)
	GetByCinema(cinemaID uuid.UUID, day *time.Time) ([]ScreeningOutput, error)
	GetAll(day *time.Time) ([]ScreeningOutput, error)
	UpdateStartTime(id uuid.UUID, newTime time.Time) error
	Delete(id uuid.UUID) error
}

type service struct {
	repo screening.Repository
}

func NewScreeningService(repo screening.Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(input CreateScreeningInput) (*ScreeningOutput, error) {
	if input.MovieID == uuid.Nil || input.RoomID == uuid.Nil || input.StartTime.IsZero() {
		return nil, errors.New("movie ID, room ID, and start time are required")
	}

	screeningID := uuid.New()
	newScreening := &screening.ScreeningInput{
		ID:        screeningID,
		MovieID:   input.MovieID,
		RoomID:    input.RoomID,
		StartTime: input.StartTime,
	}

	if err := s.repo.Create(newScreening); err != nil {
		return nil, err
	}

	createdScreening, err := s.repo.GetByID(screeningID)
	if err != nil {
		return nil, err
	}

	return toOutput(createdScreening), nil
}

func (s *service) GetByID(id uuid.UUID) (*ScreeningOutput, error) {
	sc, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return toOutput(sc), nil
}

func (s *service) GetByMovie(movieID uuid.UUID, day *time.Time) ([]ScreeningOutput, error) {
	scs, err := s.repo.GetByMovie(movieID, day)
	if err != nil {
		return nil, err
	}
	return mapToOutput(scs), nil
}

func (s *service) GetByRoom(roomID uuid.UUID, day *time.Time) ([]ScreeningOutput, error) {
	scs, err := s.repo.GetByRoom(roomID, day)
	if err != nil {
		return nil, err
	}
	return mapToOutput(scs), nil
}

func (s *service) GetByCinema(cinemaID uuid.UUID, day *time.Time) ([]ScreeningOutput, error) {
	scs, err := s.repo.GetByCinema(cinemaID, day)
	if err != nil {
		return nil, err
	}
	return mapToOutput(scs), nil
}

func (s *service) GetAll(day *time.Time) ([]ScreeningOutput, error) {
	scs, err := s.repo.GetAll(day)
	if err != nil {
		return nil, err
	}
	return mapToOutput(scs), nil
}

func (s *service) UpdateStartTime(id uuid.UUID, newTime time.Time) error {
	if id == uuid.Nil || newTime.IsZero() {
		return errors.New("invalid ID or time")
	}
	return s.repo.UpdateStartTime(id, newTime)
}

func (s *service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
