package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sn-auth/configs"
	v1 "sn-auth/internal/controller/http/v1"
	"sn-auth/internal/repo"
	"sn-auth/internal/service"
	"sn-auth/pkg/hasher"
	"sn-auth/pkg/http_server"
	"sn-auth/pkg/postgres"
	"syscall"
)

// @title           Authorization Service
// @version         1.0
// @description     This is a service for app users in system

// @contact.name   George Epishev
// @contact.email  epishcom@gmail.com

// @host      localhost:8089
// @BasePath  /

// @securityDefinitions.apikey  JWT
// @in                          header
// @name                        Authorization
// @description					JWT token

func Run(configPath string) {
	const fn = "app.Run"
	cfg, err := configs.MustLoad(configPath)
	if err != nil {
		log.Fatalf("%s: %v", fn, err)
	}

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
	deps := service.ServicesDependencies{
		Repos:    repositories,
		Hasher:   hasher.NewSHA1Hasher(cfg.Hasher.Salt),
		SignKey:  cfg.JWT.SignKey,
		TokenTTL: cfg.JWT.TokenTTL,
	}

	services := service.NewService(deps)

	log.Info("Initializing handlers and routes...")
	handler := echo.New()
	//TODO add validator
	v1.NewRouter(handler, services)

	log.Info("Starting http server...")
	log.Debugf("Server port: %d", cfg.HTTP.Port)
	httpServer := http_server.New(handler, http_server.Port(cfg.HTTP.Port))

	log.Info("Configuring graceful shutdown...")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Infof("%s: %s", fn, s.String())
	case err = <-httpServer.Notify():
		log.Error(fmt.Errorf("%s: %v", fn, err))
	}

	log.Info("Shutting down...")
	err = httpServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("%s: %v", fn, err))
	}
}
