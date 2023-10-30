package app

import (
	"github.com/sirupsen/logrus"
	"github.com/zarasfara/pet-adoption-platform/internal/config"
	"github.com/zarasfara/pet-adoption-platform/internal/delivery/http"
	"github.com/zarasfara/pet-adoption-platform/internal/repository"
	"github.com/zarasfara/pet-adoption-platform/internal/server"
	"github.com/zarasfara/pet-adoption-platform/internal/service"
)

// Run - Запуск приложения
func Run(cfg *config.Config) {
	// Инициализация БД.
	db, err := repository.NewPostgresDB(*cfg)
	if err != nil {
		logrus.Fatalf("error while connecting to database: %s", err)
	}

	// todo подключение сервисиов
	repositories := repository.NewRepository(db)

	services := service.NewService(repositories)

	handler := http.NewHandler(services)

	// Инициализация сервера и машрутов.
	srv := server.NewServer(cfg, handler.Init())

	// Запуск сервера.
	if err := srv.Serve(); err != nil {
		logrus.Fatalf("error while serve: %s", err)
	}
}
