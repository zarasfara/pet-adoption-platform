package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Authorization
	Pet
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Pet:           NewPetPostgres(db),
	}
}
