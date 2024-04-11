package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/zarasfara/pet-adoption-platform/internal/models"
)

type Pet interface {
	PetsBySortField(sortField string) ([]models.Pet, error)
}

var _ Pet = petPostgres{}

type petPostgres struct {
	db *sqlx.DB
}

func (p petPostgres) PetsBySortField(sortField string) ([]models.Pet, error) {
	query := fmt.Sprintf(`
		SELECT p.id, p.description, p.name, p.age, p.is_available, s.name AS shelter_name, b.name AS breed
		FROM %s as p
		INNER JOIN %s as s ON p.shelter_id = s.id
		INNER JOIN %s as b ON p.breed_id = b.id;
	`, petsTable, sheltersTable, breedsTable)

	if sortField != "" {
		query += fmt.Sprintf(" ORDER BY %s", sortField)
	}

	pets := make([]models.Pet, 0)

	err := p.db.Select(&pets, query)
	if err != nil {
		return nil, err
	}

	return pets, nil
}

func NewPetPostgres(db *sqlx.DB) *petPostgres {
	return &petPostgres{db: db}
}
