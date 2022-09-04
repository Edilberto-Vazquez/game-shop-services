package routes

import (
	"net/http"

	"github.com/Edilberto-Vazquez/game-shop-services/src/apigateway/drivers/http/server"
	"github.com/gin-gonic/gin"
)

func sessionRoutes(s server.Server, rg *gin.RouterGroup) {
	session := rg.Group("/session")
	services := s.Services()

	session.POST("/signup", func(ctx *gin.Context) {
		signUpModel := services.UserService.SignUpModel
		if err := ctx.ShouldBindJSON(&signUpModel); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if userId, err := services.UserService.SignUp(ctx, &signUpModel); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{"user": userId})
			return
		}
	})

	session.GET("/login", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "login")
	})
}
