package routes

import (
	"github.com/Edilberto-Vazquez/game-shop-services/src/apigateway/drivers/http/handlers"
	"github.com/Edilberto-Vazquez/game-shop-services/src/apigateway/drivers/http/server"
	"github.com/gin-gonic/gin"
)

func sessionRoutes(s server.Server, rg *gin.RouterGroup) {
	services := s.Services()
	session := rg.Group("/session")
	session.POST("/signup", handlers.SignUp(services))
	session.GET("/login", handlers.Login(services))
	session.GET("/me", handlers.Me(services))
}
