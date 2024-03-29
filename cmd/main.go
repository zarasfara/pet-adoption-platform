package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/zarasfara/pet-adoption-platform/internal/app"
	"github.com/zarasfara/pet-adoption-platform/internal/config"
)

//	@title						Pet adoption platform api
//	@version					1.0
//	@host						localhost:12001
//	@contact.email				oev2001@gmail.com
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
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
