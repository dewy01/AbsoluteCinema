package main

import (
	"absolutecinema/internal/app"
	"absolutecinema/internal/config"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		fmt.Printf("create config: %w", err)
		os.Exit(1)
	}

	ctx, cancel := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	group, _ := errgroup.WithContext(ctx)

	app, err := app.New(cfg)
	if err != nil {
		fmt.Printf("create app: %w", err)
		os.Exit(1)
	}

	fmt.Println(app.DB.Config)

	// if err := app.Start(groupCtx, group); err != nil {
	// 	fmt.Printf("Start new app failed: '%v', Aborting startup\n", err)
	// 	os.Exit(1)
	// }

	if err := group.Wait(); err != nil {
		fmt.Printf("Start new app failed: '%v', Aborting startup\n", err)
		os.Exit(1)
	}

}
