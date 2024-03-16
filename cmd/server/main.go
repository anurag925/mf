package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/anurag925/identity/app/routes"
	"github.com/anurag925/identity/cmd"
	"github.com/anurag925/identity/config"
	"github.com/anurag925/identity/core"
)

func main() {
	slog.Info("Starting the application from the main...")
	flags := cmd.ParseFlags()
	slog.Info("Application is starting on", slog.Any("environment", config.Env(config.Environment(*flags.Env))), slog.Int("port", *flags.Port))
	routes.Init()

	// Start server
	go func() {
		if err := core.Server().Start(fmt.Sprintf(":%d", *flags.Port)); err != nil && err != http.ErrServerClosed {
			slog.Error("shutting down the server")
			panic("unable to start the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()
	if err := core.Server().Shutdown(ctx); err != nil {
		slog.Error("unable to gracefully stop the server ", "error", err)
		panic("unable to gracefully stop the server")
	}
}
