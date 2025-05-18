package models

import "github.com/google/uuid"

type Actor struct {
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name   string
	Movies []Movie `gorm:"many2many:movie_actors;"`
}
