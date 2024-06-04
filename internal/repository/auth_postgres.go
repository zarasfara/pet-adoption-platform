package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/zarasfara/pet-adoption-platform/internal/models"
)

type Authorization interface {
	CreateUser(user models.AddRecordUser) error
	UserByEmail(email string) (models.User, error)
	UserByID(userID int) (models.User, error)
}

var _ Authorization = authPostgres{}

type authPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *authPostgres {
	return &authPostgres{
		db: db,
	}
}

func (r authPostgres) CreateUser(user models.AddRecordUser) error {

	query := fmt.Sprintf("INSERT INTO %s (name, email, password, preferences) VALUES ($1, $2, $3, $4)", usersTable)

	_, err := r.db.Exec(query, user.Name, user.Email, user.Password, user.Preferences)
	if err != nil {
		return err
	}

	return nil
}

func (r authPostgres) UserByEmail(email string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE email = $1", usersTable)
	err := r.db.Get(&user, query, email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r authPostgres) UserByID(userID int) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT name, email, preferences, created_at, updated_at FROM %s WHERE id = $1", usersTable)
	err := r.db.Get(&user, query, userID)
	if err != nil {
		return user, err
	}

	return user, nil
}
