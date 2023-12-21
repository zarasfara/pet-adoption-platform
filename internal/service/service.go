package service

import (
	"github.com/zarasfara/pet-adoption-platform/internal/config"
	"github.com/zarasfara/pet-adoption-platform/internal/models"
	"github.com/zarasfara/pet-adoption-platform/internal/repository"
)

type Authorization interface {
	CreateUser(user models.User) error
	GenerateToken(email, password string) (token string, err error)
	ParseToken(token string) (userId int, err error)
	GetCurrentUser(userId int) (models.User, error)
}

type Services struct {
	Authorization
}

func NewService(repos *repository.Repository, cfg *config.Config) *Services {
	return &Services{
		Authorization: NewAuthService(repos.Authorization, cfg),
	}
}