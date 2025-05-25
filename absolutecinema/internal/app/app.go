package app

import (
	"absolutecinema/internal/auth"
	"absolutecinema/internal/config"
	"absolutecinema/internal/database"
	"absolutecinema/internal/database/repository"
	"absolutecinema/internal/handlers"
	"absolutecinema/internal/openapi/gen/usergen"
	"absolutecinema/internal/service"
	"absolutecinema/pkg/log"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

type App struct {
	cfg    *config.AppConfig
	db     *database.Database
	logger *slog.Logger

	repositories *repository.Repositories
	server       *http.Server
}

func New(cfg *config.AppConfig) (*App, error) {
	logger, err := log.New(cfg.Log.Level, cfg.Log.FilePath, cfg.Mode.IsProd())
	if err != nil {
		return nil, fmt.Errorf("create logger: %w", err)
	}

	db, err := database.New(cfg.DB)
	if err != nil {
		return nil, fmt.Errorf("create db connection: %w", err)
	}

	err = db.Setup()
	if err != nil {
		return nil, fmt.Errorf("setup db: %w", err)
	}

	sessionStore := auth.NewStoreMemory()
	sessionService, err := auth.NewService(sessionStore)
	if err != nil {
		return nil, err
	}

	router := http.NewServeMux()
	repositories := repository.NewRepositories(db.Gorm())
	services := service.NewServices(repositories, sessionService)
	handlers := handlers.NewHandlers(services)

	openapiHandler := usergen.Handler(handlers.User)
	router.Handle("/users/", openapiHandler)

	const defaultTimeout = 10 * time.Second
	httpServer := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler:           router,
		ReadTimeout:       defaultTimeout,
		ReadHeaderTimeout: defaultTimeout,
		WriteTimeout:      defaultTimeout,
		IdleTimeout:       time.Minute,
	}

	return &App{
		cfg:          cfg,
		db:           db,
		logger:       logger,
		repositories: repositories,
		server:       httpServer,
	}, nil
}

func (app *App) Start(ctx context.Context, group *errgroup.Group) error {
	group.Go(func() error {
		app.logger.Info("Starting server", slog.String("address", app.server.Addr))

		if err := app.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("start server: %w", err)
		}

		app.logger.Info("Shutting down server...")
		return nil
	})

	group.Go(func() error {
		<-ctx.Done()

		app.logger.Info("Starting graceful shutdown")
		if err := app.stop(); err != nil {
			return fmt.Errorf("graceful shutdown: %w", err)
		}

		app.logger.Info("Graceful shutdown completed")
		return nil
	})

	return nil
}

func (app *App) stop() error {
	const shutdownTimeout = 10 * time.Second
	shutdownCtx, cancelFn := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancelFn()

	if err := app.server.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("gracefully stopping server: %w", err)
	}

	return nil
}
