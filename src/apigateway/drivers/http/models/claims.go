package models

import "github.com/golang-jwt/jwt/v4"

type AppClaims struct {
	UserID string `json:"userID"`
	jwt.RegisteredClaims
}
