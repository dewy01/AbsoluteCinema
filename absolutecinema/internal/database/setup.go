package database

import (
	"absolutecinema/internal/database/models"

	"gorm.io/gorm"
)

func (db *Database) Gorm() *gorm.DB {
	return db.DB
}

func (db *Database) Setup(seed bool) error {
	db.initUUID()
	if err := db.migrate(); err != nil {
		return err
	}

	if seed && db.isEmpty() {
		db.seed()
	}

	return nil
}

func (db *Database) migrate() error {
	return db.AutoMigrate(
		&models.User{},
		&models.Cinema{},
		&models.Room{},
		&models.Seat{},
		&models.Actor{},
		&models.Movie{},
		&models.Screening{},
		&models.Reservation{},
		&models.ReservedSeat{},
	)
}

func (db *Database) initUUID() {
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
}

func (db *Database) isEmpty() bool {
	var count int64
	db.Model(&models.Cinema{}).Count(&count)
	if count > 0 {
		return false
	}

	return true
}
