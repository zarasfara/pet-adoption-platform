package models

import "time"

type RefreshSession struct {
	UserId       int           `db:"user_id"`
	RefreshToken string        `db:"refresh_token"`
	ExpiresIn    int64         `db:"expires_in"`
	CreatedAt    time.Time     `db:"created_at"`
}