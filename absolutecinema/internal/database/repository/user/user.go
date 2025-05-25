package user

import (
	"absolutecinema/internal/auth"
	"absolutecinema/internal/database/models"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Role     auth.Role
	Password string
}

func ToDBUser(u *User) *models.User {
	return &models.User{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Role:     u.Role,
		Password: u.Password,
	}
}

func ToDomainUser(u *models.User) *User {
	return &User{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Role:     u.Role,
		Password: u.Password,
	}
}
