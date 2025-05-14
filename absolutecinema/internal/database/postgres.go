package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Name     string
	User     string
	Host     string
	Port     int
	Password string
}

type Database struct {
	*gorm.DB
}

func (cfg *Config) ParseConfig() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)
}

func New(cfg Config) (*Database, error) {
	db, err := gorm.Open(postgres.Open(cfg.ParseConfig()), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect with gorm: %w", err)
	}

	return &Database{DB: db}, nil
}
