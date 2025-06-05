package models

import "github.com/google/uuid"

type Movie struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title       string
	Director    string
	Description string
	PhotoPath   string
	Actors      []Actor     `gorm:"many2many:movie_actors;"`
	Screenings  []Screening `gorm:"foreignKey:MovieID"`
}
