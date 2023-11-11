package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id          int            `json:"-" db:"id"`
	Name        string         `json:"name" binding:"required"`
	Email       string         `json:"email" binding:"required"`
	Password    string         `json:"password" binding:"required" db:"password"`
	Preferences sql.NullString `json:"preferences"` // Предпочтения
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
}
