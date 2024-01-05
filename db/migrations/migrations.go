package migrate

import (
	"log/slog"

	"github.com/uptrace/bun/migrate"
)

func SqlMigrate() {
	slog.Info("Starting migration for sql...")
	if err := migrate.NewMigrations().DiscoverCaller(); err != nil {
		slog.Info("error in migrating sql", slog.Any("error", err))
		panic(err)
	}
	slog.Info("Migration for sql is done...")
}
