package service

import (
	"github.com/zarasfara/pet-adoption-platform/internal/models"
	"github.com/zarasfara/pet-adoption-platform/internal/repository"
)

type Pet interface {
	PetsBySortField(sortField string) ([]models.Pet, error)
}

var _ Pet = petService{}

type petService struct {
	repo repository.Pet
}

func (p petService) PetsBySortField(sortField string) ([]models.Pet, error) {
	pets, err := p.repo.PetsBySortField(sortField)
	if err != nil {
		return nil, err
	}

	return pets, nil
}

func NewPetService(repo repository.Pet) *petService {
	return &petService{repo: repo}
}
