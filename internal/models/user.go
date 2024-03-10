package models

import (
	"time"
)

type User struct {
	Id          int       `json:"-"`
	Name        string    `json:"name" binding:"required" example:"Eugene" extensions:"string"`
	Email       string    `json:"email" binding:"required" extensions:"string" example:"test@gmail.com"`
	Password    string    `json:"-" binding:"required" db:"password" example:"password"`
	Preferences *string   `json:"preferences" swaggertype:"string" example:"some text about my preferences"` // Предпочтения
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}
