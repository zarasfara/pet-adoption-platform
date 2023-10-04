package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zarasfara/pet-adoption-platform/internal/app"
	"github.com/zarasfara/pet-adoption-platform/internal/config"
)

func main() {
	// read config from dotenv.
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	// create config struct and fill it.
	cfg, err := config.Init(os.Getenv("APP_ENV"))
	if err != nil {
		log.Fatalf("Error while init config: %s", err)
	}

	app.Run(cfg)
}
