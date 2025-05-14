package app

import (
	"absolutecinema/internal/config"
	"absolutecinema/internal/database"
)

type App struct {
	DB *database.Database
}

func New(cfg *config.AppConfig) (*App, error) {
	db, err := database.New(cfg.DB)
	if err != nil {
		return nil, err
	}

	err = db.Setup()
	if err != nil {
		return nil, err
	}

	return &App{
		DB: db,
	}, nil
}
