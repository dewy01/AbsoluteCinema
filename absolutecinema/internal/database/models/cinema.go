package models

import "github.com/google/uuid"

type Cinema struct {
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name    string
	Address string
	Rooms   []Room
}
