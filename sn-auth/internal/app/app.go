package app

import (
	log "github.com/sirupsen/logrus"
	"sn-auth/internal/config"
	"sn-auth/internal/repo"
	"sn-auth/pkg/postgres"
)

func Run(configPath string) {
	const fn = "app.Run"
	cfg := config.MustLoad()

	SetLogrus(cfg.Log.Level)

	log.Info("Initializing postgres...")
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.MaxPoolSize))
	if err != nil {
		log.Fatal("%s: %w", fn, err)
	}
	defer pg.Close()

	log.Info("Initializing repositories...")
	repositories := repo.NewRepositories(pg)

	log.Info("Initializing services...")
	deps := service.Se
}
