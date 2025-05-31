package auth

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
)

const CookieName = "AbosluteCinemaSession"

type Key string

var ErrNotFound = errors.New("session not found")

type Store interface {
	Delete(ctx context.Context, sessionID uuid.UUID) error
	Create(ctx context.Context, session Session) error
	Get(ctx context.Context, sessionID uuid.UUID) (Session, error)
}

type SessionData struct {
	ID    uuid.UUID
	Name  string
	Email string
	Role  Role
}

type Session struct {
	id   uuid.UUID
	Data SessionData
}

func (s Session) ToCookie(expiresAt time.Time) *http.Cookie {
	return &http.Cookie{
		Name:     CookieName,
		Value:    s.id.String(),
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // TODO (for dev),
		SameSite: http.SameSiteLaxMode,
		Expires:  expiresAt,
	}
}

func NewInvalidCookie() *http.Cookie {
	return &http.Cookie{
		Name:     CookieName,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		Expires:  time.Time{},
	}
}
