package models

import (
	"database/sql"
)

type User struct {
	Name        string         `binding:"required"`
	Email       string         `binding:"required"`
	Password    string         `binding:"required"`
	Preferences sql.NullString // Предпочтения
}