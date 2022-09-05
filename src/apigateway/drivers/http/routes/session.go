package routes

import (
	"github.com/Edilberto-Vazquez/game-shop-services/src/apigateway/drivers/http/handlers"
	"github.com/Edilberto-Vazquez/game-shop-services/src/apigateway/drivers/http/server"
	"github.com/gin-gonic/gin"
)

func sessionRoutes(s server.Server, rg *gin.RouterGroup) {
	session := rg.Group("/session")
	session.POST("/signup", handlers.SignUp(s))
	session.GET("/login", handlers.Login(s))
	session.GET("/me", handlers.Me(s))
}
