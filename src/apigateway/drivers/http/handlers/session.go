package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/Edilberto-Vazquez/game-shop-services/src/apigateway/drivers/http/services"
	"github.com/Edilberto-Vazquez/game-shop-services/src/domains/shared/valueobjects"
	"github.com/Edilberto-Vazquez/game-shop-services/src/domains/user"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SignUp(s *services.Services) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cwt, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		var u user.User
		if err := ctx.ShouldBindJSON(&u); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if statusCode, err := s.UserSessionService.SignUp(cwt, u); err != nil {
			ctx.JSON(statusCode, err.Error())
			return
		} else {
			ctx.JSON(statusCode, "User created")
		}
	}
}

func Login(s *services.Services) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cwt, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		var login LoginRequest
		if err := ctx.ShouldBindJSON(&login); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if token, statusCode, err := s.UserSessionService.Login(cwt, login.Email, login.Password); err != nil {
			ctx.JSON(statusCode, err.Error())
			return
		} else {
			ctx.JSON(statusCode, token)
			return
		}
	}
}

func Me(s *services.Services) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		token := ctx.MustGet("token").(*jwt.Token)
		if claims, ok := token.Claims.(*valueobjects.AppClaims); ok && token.Valid {
			ctx.JSON(http.StatusOK, claims.UserEmail)
			return
		}
	}
}
