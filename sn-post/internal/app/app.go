package app

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"sn-auth/pkg/postgres"
	"sn-post/config"
)

func Run(pathConfig string) {
	cfg, err := config.New(pathConfig)
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	setLogger(cfg.Log.Level)

	log.Info("Initializing postgres")
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.MaxPoolSize))
	if err != nil {
		log.Fatalf("Postgres error: %v", err)
	}

	_ = pg

	handler := echo.New()
	_ = handler.Start(cfg.HTTP.Port)
}
