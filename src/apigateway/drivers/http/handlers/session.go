package handlers

import (
	"net/http"

	"github.com/Edilberto-Vazquez/game-shop-services/src/apigateway/drivers/http/server"
	"github.com/Edilberto-Vazquez/game-shop-services/src/user"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Person struct {
	UserName string `json:"userName" binding:"required"`
}

func SignUp(s server.Server) gin.HandlerFunc {
	services := s.Services()
	return func(ctx *gin.Context) {
		var signup user.Person
		if err := ctx.ShouldBindJSON(&signup); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if userId, err := services.SessionService.SignUp(ctx, &signup); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{"user": userId})
		}
	}
}

func Login(s server.Server) gin.HandlerFunc {
	services := s.Services()
	return func(ctx *gin.Context) {
		var login user.Login
		if err := ctx.ShouldBindJSON(&login); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if login, err := services.SessionService.Login(ctx, login); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{"login": login})
			return
		}
	}
}

func Me(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.MustGet("token").(*jwt.Token)
		if claims, ok := token.Claims.(*user.AppClaims); ok && token.Valid {
			ctx.JSON(http.StatusOK, claims.UserEmail)
			return
		}
	}
}
