package initializers

import (
	"database/sql"
	"log/slog"
	"strings"

	"github.com/anurag925/identity/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func InitDB() *bun.DB {
	dbUrl := strings.Split(config.Secrets().DBUrl, "//")[1]
	if config.IsDevelopment() {
		slog.Debug("the db url form the env is ", "url", dbUrl)
	}
	sqlDB, err := sql.Open("mysql", dbUrl)
	if err != nil {
		slog.Error("Error connecting to database while starting application", "error", err)
		panic("error connecting to database")
	}
	db := bun.NewDB(sqlDB, mysqldialect.New())
	if err := db.Ping(); err != nil {
		slog.Error("Error pinging to database while starting application", "error", err)
		panic("error pinging to database")
	}
	return db
}
