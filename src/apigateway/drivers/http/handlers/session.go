package handlers

import (
	"net/http"
	"time"

	"github.com/Edilberto-Vazquez/game-shop-services/src/apigateway/drivers/http/models"
	"github.com/Edilberto-Vazquez/game-shop-services/src/apigateway/drivers/http/server"
	userModels "github.com/Edilberto-Vazquez/game-shop-services/src/user/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Person struct {
	UserName string `json:"userName" binding:"required"`
}

func SignUp(s server.Server) gin.HandlerFunc {
	services := s.Services()
	return func(ctx *gin.Context) {
		signUpModel := userModels.Person{}
		if err := ctx.ShouldBindJSON(&signUpModel); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if userId, err := services.SessionService.SignUp(ctx, &signUpModel); err != nil {
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
		var userModel Person
		if err := ctx.ShouldBindJSON(&userModel); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if _, err := services.SessionService.Login(ctx, userModel.UserName); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		claims := models.AppClaims{
			UserID: userModel.UserName,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour * 24)),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(s.Config().JWTSecret))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"userName": userModel.UserName,
			"token":    tokenString,
		})
	}
}

func Me(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.MustGet("token").(*jwt.Token)
		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			ctx.JSON(http.StatusOK, claims.UserID)
			return
		}
	}
}
