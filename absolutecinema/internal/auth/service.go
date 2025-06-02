package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Service struct {
	store Store
}

func NewService(store Store) (*Service, error) {
	if store == nil {
		return nil, errors.New("store is nil")
	}

	return &Service{
		store: store,
	}, nil
}

func (s *Service) New(ctx context.Context, data SessionData) (Session, error) {
	sessionID := uuid.New()

	value := Session{
		id:   sessionID,
		Data: data,
	}

	if err := s.store.Create(ctx, value); err != nil {
		return Session{}, fmt.Errorf("create session: %w", err)
	}

	return value, nil
}

func (s *Service) Get(ctx context.Context, sessionID uuid.UUID) (Session, error) {
	value, err := s.store.Get(ctx, sessionID)
	if err != nil {
		return Session{}, fmt.Errorf("get session: %w", err)
	}

	return value, nil
}

func (s *Service) Delete(ctx context.Context, session Session) error {
	if err := s.store.Delete(ctx, session.id); err != nil {
		return fmt.Errorf("delete session: %w", err)
	}

	return nil
}
