package cinema

import (
	"absolutecinema/internal/database/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(c *Cinema) error
	GetByID(id uuid.UUID) (*Cinema, error)
	GetAll() ([]Cinema, error)
	Update(c *Cinema) error
	Delete(id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewCinemaRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(c *Cinema) error {
	dbC := ToDBCinema(c)
	if dbC.ID == uuid.Nil {
		dbC.ID = uuid.New()
	}
	return r.db.Create(dbC).Error
}

func (r *repository) GetByID(id uuid.UUID) (*Cinema, error) {
	var dbC models.Cinema
	if err := r.db.Preload("Rooms").First(&dbC, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return ToDomainCinema(&dbC), nil
}

func (r *repository) GetAll() ([]Cinema, error) {
	var dbCinemas []models.Cinema
	if err := r.db.Preload("Rooms").Find(&dbCinemas).Error; err != nil {
		return nil, err
	}
	cinemas := make([]Cinema, len(dbCinemas))
	for i, c := range dbCinemas {
		cinemas[i] = *ToDomainCinema(&c)
	}
	return cinemas, nil
}

func (r *repository) Update(c *Cinema) error {
	dbC := ToDBCinema(c)
	return r.db.Model(&models.Cinema{}).Where("id = ?", c.ID).Updates(dbC).Error
}

func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Cinema{}, "id = ?", id).Error
}
