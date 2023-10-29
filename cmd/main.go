package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/zarasfara/pet-adoption-platform/internal/app"
	"github.com/zarasfara/pet-adoption-platform/internal/config"
)

func main() {
	// Загрузка из dotenv.
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading .env file: %s", err.Error())
	}

	// Создание структуры конфигурации и заполнение.
	cfg, err := config.Init(os.Getenv("APP_ENV"))
	if err != nil {
		logrus.Fatalf("Error while init config: %s", err)
	}

	app.Run(cfg)
}
