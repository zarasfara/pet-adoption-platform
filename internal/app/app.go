package app

import (
	"github.com/sirupsen/logrus"
	"github.com/zarasfara/pet-adoption-platform/internal/config"
	"github.com/zarasfara/pet-adoption-platform/internal/delivery/http"
	"github.com/zarasfara/pet-adoption-platform/internal/repository"
	"github.com/zarasfara/pet-adoption-platform/internal/server"
	"github.com/zarasfara/pet-adoption-platform/internal/service"
)

// Run runs the application
func Run(cfg *config.Config) {
	db, err := repository.NewPostgresDB(*cfg)
	if err != nil {
		logrus.Fatalf("error while connecting to database: %s", err)
	}

	repositories := repository.NewRepository(db)

	services := service.NewService(repositories, cfg)

	handler := http.NewHandler(services, cfg.HTTP.AppUrl, cfg.JWT.RefreshTokenTTL)

	srv := server.NewServer(cfg, handler.Init())

	if err := srv.Serve(); err != nil {
		logrus.Fatalf("error while serve: %s", err)
	}
}
