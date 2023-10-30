package service

import (
	"github.com/zarasfara/pet-adoption-platform/internal/models"
	"github.com/zarasfara/pet-adoption-platform/internal/repository"
)

type Authorization interface {
	CreateUser(user models.User) error
}

type Services struct {
	Authorization
}

func NewService(repos *repository.Repository) *Services {
	return &Services{
		Authorization: NewAuthService(repos.Authorization),
	}
}