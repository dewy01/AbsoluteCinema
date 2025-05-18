package actor

import (
	"absolutecinema/internal/database/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(actor *Actor) error
	GetByID(id uuid.UUID) (*Actor, error)
	GetAll() ([]Actor, error)
	Update(actor *Actor) error
	Delete(id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewActorRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(actor *Actor) error {
	dbA := ToDBActor(actor)
	if dbA.ID == uuid.Nil {
		dbA.ID = uuid.New()
	}
	return r.db.Create(dbA).Error
}

func (r *repository) GetByID(id uuid.UUID) (*Actor, error) {
	var dbA models.Actor
	if err := r.db.Preload("Movies").First(&dbA, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return ToDomainActor(&dbA), nil
}

func (r *repository) GetAll() ([]Actor, error) {
	var dbActors []models.Actor
	if err := r.db.Preload("Movies").Find(&dbActors).Error; err != nil {
		return nil, err
	}
	actors := make([]Actor, len(dbActors))
	for i, a := range dbActors {
		actors[i] = *ToDomainActor(&a)
	}
	return actors, nil
}

func (r *repository) Update(actor *Actor) error {
	dbA := ToDBActor(actor)
	return r.db.Model(&models.Actor{}).Where("id = ?", actor.ID).Updates(dbA).Error
}

func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Actor{}, "id = ?", id).Error
}
