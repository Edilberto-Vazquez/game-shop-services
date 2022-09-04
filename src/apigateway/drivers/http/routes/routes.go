package routes

import (
	"github.com/Edilberto-Vazquez/game-shop-services/src/apigateway/drivers/http/server"
	"github.com/gin-gonic/gin"
)

func GetRoutes(s server.Server, r *gin.Engine) {
	v1 := r.Group("/api/v1")
	sessionRoutes(s, v1)
}