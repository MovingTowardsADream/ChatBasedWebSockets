package app

import (
	"ChatBasedWebSockets/internal/config"
	v1 "ChatBasedWebSockets/internal/controller/http/v1"
	"ChatBasedWebSockets/internal/repository/postgresdb"
	"ChatBasedWebSockets/internal/usecase"
	"ChatBasedWebSockets/pkg/hasher"
	"ChatBasedWebSockets/pkg/httpserver"
	"ChatBasedWebSockets/pkg/postgres"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type App struct {
	HTTPServer *httpserver.Server
	DB         *postgres.Postgres
}

func New(log *slog.Logger, cfg *config.Config) *App {

	// Connect postgresdb db
	pg, err := postgres.NewPostgresDB(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		panic("app - Run - postgresdb.NewPostgresDB: " + err.Error())
	}

	repoAuth := postgresdb.NewAuthRepo(pg)

	passHasher := hasher.NewSHA1Hasher("salt")

	uscsAuth := usecase.NewAuthUseCase(log, repoAuth, "signKey", passHasher, cfg.TokenTTL)

	handler := gin.New()
	v1.NewRouter(handler, log, uscsAuth)
	httpServer := httpserver.New(log, handler, httpserver.Port(cfg.HTTP.Port), httpserver.WriteTimeout(cfg.HTTP.Timeout))

	return &App{
		HTTPServer: httpServer,
		DB:         pg,
	}
}
