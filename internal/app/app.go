package app

import (
	"ChatBasedWebSockets/internal/config"
	v1 "ChatBasedWebSockets/internal/controller/http/v1"
	"ChatBasedWebSockets/internal/controller/http/v1/ws"
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
	repoUsers := postgresdb.NewUserRepo(pg)

	passHasher := hasher.NewSHA1Hasher(cfg.Salt)

	uscsAuth := usecase.NewAuthUseCase(log, repoAuth, cfg.SignKey, passHasher, cfg.TokenTTL)
	uscsUsers := usecase.NewUsersUseCase(log, repoUsers)

	gin.SetMode(cfg.GinMode)

	handler := gin.New()

	manager := ws.NewManager()

	v1.NewRouter(handler, log, uscsAuth, uscsUsers, manager)
	httpServer := httpserver.New(log, handler, httpserver.Port(cfg.HTTP.Port), httpserver.WriteTimeout(cfg.HTTP.Timeout))

	return &App{
		HTTPServer: httpServer,
		DB:         pg,
	}
}
