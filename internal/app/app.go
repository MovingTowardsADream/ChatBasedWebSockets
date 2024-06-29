package app

import (
	"ChatBasedWebSockets/internal/config"
	"ChatBasedWebSockets/pkg/httpserver"
	"ChatBasedWebSockets/pkg/postgres"
	"log/slog"
)

type App struct {
	HTTPServer *httpserver.Server
	PG         *postgres.Postgres
}

func New(log *slog.Logger, cfg *config.Config) *App {

	// Connect postgres db
	pg, err := postgres.NewPostgresDB(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		panic("app - Run - postgres.NewPostgresDB: " + err.Error())
	}
}
