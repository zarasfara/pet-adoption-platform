package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/zarasfara/pet-adoption-platform/internal/models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (r AuthPostgres) CreateUser(user models.User) error {

	query := fmt.Sprintf("INSERT INTO %s (name, email, password) VALUES ($1, $2, $3)", usersTable)

	_, err := r.db.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r AuthPostgres) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE email = $1", usersTable)
	err := r.db.Get(&user, query, email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r AuthPostgres) GetUserByID(userID int) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT name, email, preferences, created_at, updated_at FROM %s WHERE id = $1", usersTable)
	err := r.db.Get(&user, query, userID)
	if err != nil {
		return user, err
	}

	return user, nil
}
