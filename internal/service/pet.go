package service

import (
	"github.com/zarasfara/pet-adoption-platform/internal/models"
	"github.com/zarasfara/pet-adoption-platform/internal/repository"
)

var _ Pet = PetService{}

type PetService struct {
	repo repository.Pet
}

func (p PetService) GetAll(sortField string) ([]models.Pet, error) {
	pets, err := p.repo.GetAll(sortField)
	if err != nil {
		return nil, err
	}

	return pets, nil
}

func NewPetService(repo repository.Pet) *PetService {
	return &PetService{repo: repo}
}
