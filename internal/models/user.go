package models

import (
	"database/sql"
)

type User struct {
	Name        string         `json:"name" binding:"required"`
	Email       string         `json:"email" binding:"required"`
	Password    string         `json:"password" binding:"required"`
	Preferences sql.NullString `json:"preferences"` // Предпочтения
}
