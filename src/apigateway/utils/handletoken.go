package utils

import (
	"strings"

	"github.com/Edilberto-Vazquez/game-shop-services/src/apigateway/drivers/http/server"
	"github.com/Edilberto-Vazquez/game-shop-services/src/user"
	"github.com/golang-jwt/jwt/v4"
)

func RouteNeedToken(route string, NO_AUTH_NEEDED []string) bool {
	for _, p := range NO_AUTH_NEEDED {
		if strings.Contains(route, p) {
			return false
		}
	}
	return true
}

func ProcessToken(authorization string, s server.Server) (*jwt.Token, error) {
	tokenString := strings.TrimSpace(authorization)
	token, err := jwt.ParseWithClaims(tokenString, &user.AppClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.Services().SessionService.JWTSecret), nil
	})
	return token, err
}
