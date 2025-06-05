package config

import (
	"absolutecinema/internal/database"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

type AppConfig struct {
	Mode   Environment
	Log    LogConfig
	Server ServerConfig
	Seed   bool
	DB     database.Config
}

type Environment string

const (
	ModeProd Environment = "prod"
	ModeDev  Environment = "dev"
)

type LogConfig struct {
	Level    slog.Level
	FilePath string
}

type ServerConfig struct {
	Host string
	Port int
}

func New() (*AppConfig, error) {
	err := LoadEnv()
	if err != nil {
		return nil, err
	}

	dbport, err := strconv.Atoi(GetEnv("AC_DATABASE_PORT", "5432"))
	if err != nil {
		return nil, err
	}

	serverport, err := strconv.Atoi(GetEnv("AC_SERVER_PORT", "8080"))
	if err != nil {
		return nil, err
	}

	seed, err := strconv.ParseBool(GetEnv("AC_SEED", "true"))
	if err != nil {
		return nil, err
	}

	modeStr := GetEnv("AC_MODE", "dev")
	mode, err := ParseMode(modeStr)
	if err != nil {
		return nil, err
	}

	level, err := ParseLogLevel(GetEnv("AC_LOG_LEVEL", "info"))
	if err != nil {
		return nil, err
	}

	cfg := &AppConfig{
		Log: LogConfig{
			Level:    level,
			FilePath: GetEnv("AC_LOG_FILE_PATH", "."),
		},
		Mode: mode,
		DB: database.Config{
			Name:     GetEnv("AC_DATABASE_NAME", ""),
			User:     GetEnv("AC_DATABASE_USER", ""),
			Password: GetEnv("AC_DATABASE_PASSWORD", ""),
			Host:     GetEnv("AC_DATABASE_HOST", "localhost"),
			Port:     dbport,
		},
		Server: ServerConfig{
			Host: GetEnv("AC_SERVER_HOST", ""),
			Port: serverport,
		},
		Seed: seed,
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

func ParseLogLevel(s string) (slog.Level, error) {
	var level slog.Level
	if err := level.UnmarshalText([]byte(s)); err != nil {
		return level, err
	}

	return level, nil
}
