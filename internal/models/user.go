package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string         `gorm:"size:50;not null" binding:"required"`
	Email       string         `gorm:"unique;not null" binding:"required"`
	Password    string         `gorm:"unique;not null" binding:"required"`
	Preferences sql.NullString // Предпочтения
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.CreatedAt = time.Now()
	return nil
}
