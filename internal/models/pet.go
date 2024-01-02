package models

import "database/sql"

type Pet struct {
	ID          int            `json:"-"`
	Description sql.NullString `json:"description,omitempty"`
	Name        string         `json:"name"`
	Age         int            `json:"age"`
	IsAvailable bool           `json:"isAvailable"`
	ShelterName string         `json:"shelterName" db:"shelter_name"`
	Breed       string         `json:"breedId" db:"breed"`
}

type AddRecordPet struct {
	Name        string         `json:"name" binding:"required,alpha"`
	Description sql.NullString `json:"description" binding:"required,alpha"`
	Age         int            `json:"age" binding:"required,number"`
	IsAvailable bool           `json:"isAvailable" binding:"required,boolean"`
	ShelterId   int            `json:"shelterId" binding:"required,number"`
	BreedId     int            `json:"breedId" binding:"required,number"`
}
