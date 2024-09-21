package app

import (
	"github.com/labstack/echo/v4"
	"github.com/magmaheat/social-network/sn-auth/pkg/postgres"
	"github.com/magmaheat/social-network/sn-post/configs"
	log "github.com/sirupsen/logrus"
)

func Run(pathConfig string) {
	cfg, err := configs.New("configs/local.yaml")
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

	if err := handler.Start(":" + cfg.HTTP.Port); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
