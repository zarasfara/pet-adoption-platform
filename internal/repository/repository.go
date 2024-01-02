package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/zarasfara/pet-adoption-platform/internal/models"
)

type Authorization interface {
	CreateUser(user models.User) error
	GetUserByEmail(email string) (models.User, error)
	GetUserByID(userID int) (models.User, error)
}

type Pet interface {
	GetAll(sortField string) ([]models.Pet, error)
}

type Repository struct {
	Authorization
	Pet
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Pet:           NewPet(db),
	}
}
