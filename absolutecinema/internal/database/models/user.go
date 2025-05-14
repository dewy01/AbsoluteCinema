package models

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name         string
	Email        string `gorm:"unique"`
	Password     string
	Reservations []Reservation
}
