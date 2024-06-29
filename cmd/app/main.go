package main

import (
	"ChatBasedWebSockets/internal/app"
	"ChatBasedWebSockets/internal/config"
	"ChatBasedWebSockets/pkg/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Log.Level)

	application := app.New(log, cfg)

	_ = application
}
