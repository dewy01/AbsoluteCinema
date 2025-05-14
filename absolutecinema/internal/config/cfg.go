package config

import (
	"absolutecinema/internal/database"
	"fmt"
	"strconv"
	"strings"
)

type AppConfig struct {
	Mode Environment
	Log  LogConfig
	DB   database.Config
}

type Environment string

const (
	ModeProd Environment = "prod"
	ModeDev  Environment = "dev"
)

type LogConfig struct {
	Level    string
	FilePath string
}

func New() (*AppConfig, error) {
	err := LoadEnv()
	if err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(GetEnv("AC_DATABASE_PORT", "5432"))
	if err != nil {
		return nil, err
	}

	modeStr := GetEnv("AC_MODE", "dev")
	mode, err := ParseMode(modeStr)
	if err != nil {
		return nil, err
	}

	cfg := &AppConfig{
		Log: LogConfig{
			Level:    GetEnv("AC_LOG_LEVEL", "info"),
			FilePath: GetEnv("AC_LOG_FILE_PATH", "."),
		},
		Mode: mode,
		DB: database.Config{
			Name:     GetEnv("AC_DATABASE_NAME", ""),
			User:     GetEnv("AC_DATABASE_USER", ""),
			Password: GetEnv("AC_DATABASE_PASSWORD", ""),
			Host:     GetEnv("AC_DATABASE_HOST", "localhost"),
			Port:     port,
		},
	}

	return cfg, nil
}

func ParseMode(s string) (Environment, error) {
	switch strings.ToLower(s) {
	case "prod":
		return ModeProd, nil
	case "dev":
		return ModeDev, nil
	default:
		return "", fmt.Errorf("invalid mode: %s (expected 'prod' or 'dev')", s)
	}
}

func (e Environment) IsProd() bool {
	return e == ModeProd
}

func (e Environment) IsDev() bool {
	return e == ModeDev
}
