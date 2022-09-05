package models

import "github.com/google/uuid"

type Person struct {
	ID        uuid.UUID `json:"id"`
	UserName  string    `json:"userName" validate:"required" binding:"required"`
	Email     string    `json:"email" validate:"required" binding:"required"`
	CountryId string    `json:"countryId" validate:"required" binding:"required"`
	Salt      string    `json:"salt" validate:"required" binding:"required"`
	Hash      string    `json:"hash" validate:"required" binding:"required"`
}
