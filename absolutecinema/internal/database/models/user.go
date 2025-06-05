package models

import (
	"absolutecinema/internal/auth"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name         string
	Email        string `gorm:"unique"`
	Role         auth.Role
	Password     string
	Reservations []Reservation `gorm:"foreignKey:UserID"`
}
