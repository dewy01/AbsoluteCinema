package models

import "github.com/google/uuid"

type Seat struct {
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Row    string
	Number int
	RoomID uuid.UUID `gorm:"type:uuid"`
	Room   Room      `gorm:"foreignKey:RoomID"`
}
