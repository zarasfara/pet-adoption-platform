package app

import (
	"log"

	"github.com/zarasfara/pet-adoption-platform/internal/config"
	v1 "github.com/zarasfara/pet-adoption-platform/internal/delivery/http"
	"github.com/zarasfara/pet-adoption-platform/internal/server"
)

// Run - Запуск приложения
func Run(cfg *config.Config) {

	srv := server.NewServer(cfg, v1.InitRoutes())

	// init database...

	//init handlers, repositories, services...

	if err := srv.Serve(); err != nil {
		log.Fatalf("error while serve: %s", err)
	}
}
