package app

import (
	"log"

	"github.com/zarasfara/pet-adoption-platform/internal/config"
	"github.com/zarasfara/pet-adoption-platform/internal/delivery/http"
	"github.com/zarasfara/pet-adoption-platform/internal/server"
)

// Run - Запуск приложения
func Run(cfg *config.Config) {
	// чтения из dotenv
	handler := http.NewHandler()

	// Инициализация севера и машрутов
	srv := server.NewServer(cfg, handler.Init())

	// инициализация БД...
	_, err := cfg.DB.NewConnection()
	if err != nil {
		log.Fatalf("error while connecting to database: %s", err)
	}
	// init handlers, repositories, services...

	// Старт сервера
	if err := srv.Serve(); err != nil {
		log.Fatalf("error while serve: %s", err)
	}
}
