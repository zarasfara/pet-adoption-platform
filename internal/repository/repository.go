package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/zarasfara/pet-adoption-platform/internal/models"
)

type Authorization interface {
	CreateUser(user models.User) error
	GetUser(email string) (models.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
