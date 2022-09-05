package handlers

import (
	"net/http"

	"github.com/Edilberto-Vazquez/game-shop-services/src/apigateway/drivers/http/server"
	"github.com/Edilberto-Vazquez/game-shop-services/src/user/models"
	"github.com/gin-gonic/gin"
)

func SignUp(s server.Server) gin.HandlerFunc {
	services := s.Services()
	return func(ctx *gin.Context) {
		signUpModel := models.Person{}
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
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "login")
	}
}
