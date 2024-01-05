package core

import (
	"database/sql"
	"log/slog"
	"os"

	"github.com/anurag925/mf/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	slogecho "github.com/samber/slog-echo"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

type App struct {
	config config.Config
	server *echo.Echo
	db     *bun.DB
}

func New() {

}

func (s *App) initServer() {
	server := echo.New()
	if config.Env() == config.Production {
		server.Logger.SetLevel(log.INFO)
		server.HideBanner = true
		server.HidePort = true
	} else {
		server.Debug = true
		server.Logger.SetLevel(log.DEBUG)
	}
	server.Use(middleware.Recover())
	server.Use(middleware.CORS())
	server.Use(slogecho.New(slog.New(slog.NewTextHandler(os.Stdout, nil))))
	// server.Use(middleware.BodyDump(func(ctx echo.Context, b1, b2 []byte) {
	// 	logger.Info(ctx.Get("context").(context.Context), "request", "method", ctx.Request().Method, "uri", ctx.Request().RequestURI, "body", b1)
	// 	logger.Info(ctx.Get("context").(context.Context), "response", "method", ctx.Request().Method, "uri", ctx.Request().RequestURI, "body", b2)
	// }))
	s.server = server
}

func (s *App) initDB() error {
	sqldb, err := sql.Open("mysql", s.config.DBUrl)
	if err != nil {
		return err
	}
	s.db = bun.NewDB(sqldb, mysqldialect.New())
	if err := s.db.Ping(); err != nil {
		return err
	}
	return nil
}
