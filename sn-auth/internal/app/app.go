package app

import (
	log "github.com/sirupsen/logrus"
	"sn-auth/configs"
	"sn-auth/internal/repo"
	"sn-auth/internal/service"
	"sn-auth/pkg/hasher"
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

	log.Info("Initializing services...")
	deps := service.ServicesDependencies{
		Repos:    repositories,
		Hasher:   hasher.NewBCRYPTHasher(),
		SignKey:  cfg.SignKey,
		TokenTTL: cfg.TokenTTL,
	}
	services := service.NewServices(deps)

	_ = services
	// TODO start server
}
