package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/zarasfara/pet-adoption-platform/internal/config"
)

const (
	usersTable = "users"
)

func NewPostgresDB(cfg config.Config) (*sqlx.DB, error) {
	dsn := "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s"
	db, err := sqlx.Connect("postgres", fmt.Sprintf(dsn, cfg.DB.Host, cfg.DB.Username, cfg.DB.Password, cfg.DB.Database, cfg.DB.Port, cfg.DB.SSLMode))
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
