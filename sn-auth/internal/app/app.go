package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"os"
	"sn-auth/configs"
	v1 "sn-auth/internal/controller/htttp/v1"
	"sn-auth/internal/repo"
	"sn-auth/internal/service"
	"sn-auth/pkg/hasher"
	"sn-auth/pkg/httpserver"
	"sn-auth/pkg/postgres"
	"sn-auth/pkg/validator"
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

	log.Info("Initializing handlers and routes...")
	handler := echo.New()

	handler.Validator = validator.NewCustomValidator()
	v1.NewRouter(handler, services)

	log.Info("Starting http server...")
	log.Debugf("Server port: %d", cfg.HTTP.Port)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	log.Info("Configuring graceful shutdown...")
	interrupt := make(chan os.Signal, 1)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Info(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	log.Info("Shutting down...")
	err = httpServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
