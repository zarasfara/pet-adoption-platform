package repository

import (
	"fmt"

	"github.com/zarasfara/pet-adoption-platform/internal/config"
	"github.com/zarasfara/pet-adoption-platform/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	usersTable = "users"
)

func NewPostgresDB(cfg config.Config) (*gorm.DB, error) {
	dsn := "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s"
	db, err := gorm.Open(postgres.Open(fmt.Sprintf(dsn, cfg.DB.Host, cfg.DB.Username, cfg.DB.Password, cfg.DB.Database, cfg.DB.Port, cfg.DB.SSLMode)))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{})

	return db, nil
}
