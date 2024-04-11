package service

import (
	"github.com/zarasfara/pet-adoption-platform/internal/config"
	"github.com/zarasfara/pet-adoption-platform/internal/repository"
)

type Services struct {
	Authorization
	Pet
}

func NewService(repos *repository.Repository, cfg *config.Config) *Services {
	return &Services{
		Authorization: NewAuthService(repos.Authorization, cfg),
		Pet:           NewPetService(repos.Pet),
	}
}
