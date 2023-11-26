package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id          int            `json:"-" db:"id"`
	Name        string         `json:"name" binding:"required" example:"Eugene"`
	Email       string         `json:"email" binding:"required" extensions:"string" example:"test@gmail.com"`
	Password    string         `json:"password" binding:"required" db:"password" example:"password"`
	Preferences sql.NullString `json:"preferences" swaggertype:"string" example:"some text about my preferences"` // Предпочтения
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
}
