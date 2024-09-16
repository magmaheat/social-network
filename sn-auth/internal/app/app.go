package app

import (
	log "github.com/sirupsen/logrus"
	"sn-auth/configs"
	"sn-auth/internal/repo"
	"sn-auth/pkg/postgres"
)

func Run(configPath string) {
	cfg, err := configs.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	SetLogger(cfg.Log.Level)

	log.Info("Initializing postgres...")
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.MaxPoolSize))
	if err != nil {
		log.Fatalf("app - Run - pgdb.NewServices: %v", err)
	}
	defer pg.Close()

	log.Info("Initializing repositories...")
	repositories := repo.NewRepositories(pg)

	_ = repositories

	log.Info("Initializing services...")

	// TODO init service

	// TODO start server
}
