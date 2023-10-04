package app

import (
	"log"

	"github.com/zarasfara/pet-adoption-platform/internal/config"
	"github.com/zarasfara/pet-adoption-platform/internal/delivery/http"
	"github.com/zarasfara/pet-adoption-platform/internal/server"
)

// Run - Запуск приложения
func Run(cfg *config.Config) {

	handler := http.NewHandler()

	srv := server.NewServer(cfg, handler.Init())

	// init database...

	//init handlers, repositories, services...

	if err := srv.Serve(); err != nil {
		log.Fatalf("error while serve: %s", err)
	}
}
