package user_service

import (
	"absolutecinema/internal/auth"

	"github.com/google/uuid"
)

type CreateUserInput struct {
	Name            string
	Email           string
	Password        string
	ConfirmPassword string
}

type LoginInput struct {
	Email    string
	Password string
}

type UserOutput struct {
	ID    uuid.UUID
	Name  string
	Email string
	Role  auth.Role
}
