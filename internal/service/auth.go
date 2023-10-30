package service

import (
	"github.com/zarasfara/pet-adoption-platform/internal/models"
	"github.com/zarasfara/pet-adoption-platform/internal/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s AuthService) CreateUser(user models.User) error {
	return s.repo.CreateUser(user)
}
