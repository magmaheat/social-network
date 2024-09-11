package app

import (
	log "github.com/sirupsen/logrus"
	"sn-auth/internal/config"
)

func Run(configPath string) {
	cfg := config.MustLoad()

	SetLogrus(cfg.Log.Level)

	log.Info("Initializing postgres...")
}
