package valueobjects

import "github.com/golang-jwt/jwt/v4"

type AppClaims struct {
	UserEmail string `json:"userEmail"`
	jwt.RegisteredClaims
}
