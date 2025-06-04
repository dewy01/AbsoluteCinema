package app

import (
	"absolutecinema/internal/auth"
	"absolutecinema/internal/config"
	"absolutecinema/internal/database"
	"absolutecinema/internal/database/repository"
	"absolutecinema/internal/handlers"
	"absolutecinema/internal/openapi/gen/actorgen"
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

	mux := http.NewServeMux()

	repositories := repository.NewRepositories(db.Gorm())
	services := service.NewServices(repositories, sessionService)
	handlers := handlers.NewHandlers(services)

	// TODO MIDDLEWARE
	// openapiHandler := usergen.HandlerWithOptions(handlers.User, usergen.StdHTTPServerOptions{
	// 	Middlewares: []usergen.MiddlewareFunc{
	// 		middleware.AuthenticationMiddleware(sessionService),
	// 	},
	// })

	mux.Handle("/users/", usergen.Handler(handlers.User))
	mux.Handle("/actors/", actorgen.Handler(handlers.Actor))

	const defaultTimeout = 10 * time.Second
	httpServer := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler:           withCORS(mux),
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

func withCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        w.Header().Set("Access-Control-Allow-Credentials", "true")

        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusNoContent)
            return
        }

        next.ServeHTTP(w, r)
    })
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
