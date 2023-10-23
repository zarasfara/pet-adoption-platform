package app

import (
	"log"

	"github.com/zarasfara/pet-adoption-platform/internal/config"
	"github.com/zarasfara/pet-adoption-platform/internal/delivery/http"
	"github.com/zarasfara/pet-adoption-platform/internal/repository"
	"github.com/zarasfara/pet-adoption-platform/internal/server"
)

// Run - Запуск приложения
func Run(cfg *config.Config) {
	// Чтение из dotenv.
	handler := http.NewHandler()

	// Инициализация сервера и машрутов.
	srv := server.NewServer(cfg, handler.Init())

	// Инициализация БД.
	db, err := repository.NewPostgresDB(*cfg)
	if err != nil {
		log.Fatalf("error while connecting to database: %s", err)
	}

	// todo подключение сервисиов
	_ = repository.NewRepository(db)

	// Запуск сервера.
	if err := srv.Serve(); err != nil {
		log.Fatalf("error while serve: %s", err)
	}
}
