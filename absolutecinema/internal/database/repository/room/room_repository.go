package room

import (
	"absolutecinema/internal/database/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(room *Room) error
	GetByID(id uuid.UUID) (*Room, error)
	GetByCinemaID(cinemaID uuid.UUID) ([]Room, error)
	Update(room *Room) error
	Delete(id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(room *Room) error {
	dbRoom := ToDBRoom(room)
	if dbRoom.ID == uuid.Nil {
		dbRoom.ID = uuid.New()
	}
	for i := range dbRoom.Seats {
		if dbRoom.Seats[i].ID == uuid.Nil {
			dbRoom.Seats[i].ID = uuid.New()
		}
	}
	return r.db.Create(dbRoom).Error
}

func (r *repository) GetByID(id uuid.UUID) (*Room, error) {
	var dbRoom models.Room
	if err := r.db.Preload("Seats").First(&dbRoom, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return ToDomainRoom(&dbRoom), nil
}

func (r *repository) GetByCinemaID(cinemaID uuid.UUID) ([]Room, error) {
	var dbRooms []models.Room
	if err := r.db.Preload("Seats").Where("cinema_id = ?", cinemaID).Find(&dbRooms).Error; err != nil {
		return nil, err
	}

	rooms := make([]Room, len(dbRooms))
	for i, dbRoom := range dbRooms {
		rooms[i] = *ToDomainRoom(&dbRoom)
	}
	return rooms, nil
}

func (r *repository) Update(room *Room) error {
	return r.db.Model(&models.Room{}).Where("id = ?", room.ID).Updates(map[string]interface{}{
		"name":      room.Name,
		"cinema_id": room.CinemaID,
	}).Error
}

func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Room{}, "id = ?", id).Error
}
