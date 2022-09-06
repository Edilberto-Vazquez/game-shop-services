package models

import "time"

type Paranoid struct {
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

type Person struct {
	ID        string `json:"id"`
	UserName  string `json:"userName" validate:"required"`
	Email     string `json:"email" validate:"required"`
	CountryId string `json:"countryId" validate:"required"`
	Password  string `json:"password" validate:"required"`
}
