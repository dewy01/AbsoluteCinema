package user

import (
	"absolutecinema/internal/database/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *User) error
	GetByEmail(email string) (*User, error)
	GetByID(id uuid.UUID) (*User, error)
	Update(user *User) error
	Delete(id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(user *User) error {
	dbUser := ToDBUser(user)
	if dbUser.ID == uuid.Nil {
		dbUser.ID = uuid.New()
	}
	return r.db.Create(dbUser).Error
}

func (r *repository) GetByEmail(email string) (*User, error) {
	var dbUser models.User
	if err := r.db.Where("email = ?", email).First(&dbUser).Error; err != nil {
		return nil, err
	}
	return ToDomainUser(&dbUser), nil
}

func (r *repository) GetByID(id uuid.UUID) (*User, error) {
	var dbUser models.User
	if err := r.db.First(&dbUser, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return ToDomainUser(&dbUser), nil
}

func (r *repository) Update(user *User) error {
	dbUser := ToDBUser(user)
	return r.db.Model(&models.User{}).Where("id = ?", user.ID).Updates(map[string]interface{}{
		"name":     dbUser.Name,
		"email":    dbUser.Email,
		"password": dbUser.Password,
	}).Error
}

func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.User{}, "id = ?", id).Error
}
