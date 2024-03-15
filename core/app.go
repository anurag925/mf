package core

import (
	"log/slog"
	"sync"

	"github.com/anurag925/identity/config"
	"github.com/anurag925/identity/core/initializers"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type application struct {
	logger *slog.Logger
	server *echo.Echo
	db     *bun.DB
}

var (
	app     application
	appOnce sync.Once
)

func App() application {
	appOnce.Do(func() {
		slog.Info("The app env is ", "env", config.Env())
		if config.IsDevelopment() {
			slog.Info("The app is running in development mode")
			slog.Info("The app", "settings", config.Settings())
			slog.Info("The app", "secrets", config.Secrets())
		}
		app.logger = initializers.InitLogger()
		app.db = initializers.InitDB()
		app.server = initializers.InitServer()
	})
	return app
}

func Script() application {
	appOnce.Do(func() {
		slog.Info("The app env is ", "env", config.Env())
		if config.IsDevelopment() {
			slog.Info("The app is running in development mode")
			slog.Info("The app", "settings", config.Settings())
			slog.Info("The app", "secrets", config.Secrets())
		}
		app.logger = initializers.InitLogger()
		app.db = initializers.InitDB()
	})
	return app
}

func Server() *echo.Echo {
	return App().server
}

func DB() *bun.DB {
	return App().db
}
