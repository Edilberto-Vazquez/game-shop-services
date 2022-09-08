package user

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AppClaims struct {
	UserEmail string `json:"userEmail"`
	jwt.RegisteredClaims
}

type Paranoid struct {
	CreatedAt time.Time `json:"createdAt" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deletedAt" bson:"deleted_at,omitempty"`
}

type Person struct {
	ID        string `json:"id" bson:"_id"`
	UserName  string `json:"userName" validate:"required" bson:"user_name"`
	Email     string `json:"email" validate:"required" bson:"email"`
	CountryId string `json:"countryId" validate:"required" bson:"country_id"`
	Password  string `json:"password" validate:"required" bson:"password"`
}

type Login struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Response interface {
	StatusCode() int
	Error() error
}

type SignUpResponse struct {
	Code int   `json:"statusCode"`
	Err  error `json:"error"`
}

func (sr SignUpResponse) StatusCode() int {
	return sr.Code
}

func (sr SignUpResponse) Error() error {
	return sr.Err
}

type LoginResponse struct {
	Token string `json:"token"`
	Code  int    `json:"statusCode"`
	Err   error  `json:"error"`
}

func (lr LoginResponse) StatusCode() int {
	return lr.Code
}

func (lr LoginResponse) Error() error {
	return lr.Err
}
