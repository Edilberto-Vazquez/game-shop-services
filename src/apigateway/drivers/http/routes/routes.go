package routes

import (
	"time"

	"github.com/Edilberto-Vazquez/game-shop-services/src/apigateway/drivers/http/middleware"
	"github.com/Edilberto-Vazquez/game-shop-services/src/apigateway/drivers/http/server"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRoutes(s server.Server, r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.Use(middleware.CheckAuthMiddleware(s))
	v1.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowMethods:     []string{"GET", "POST"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Authorization"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	sessionRoutes(s, v1)
}
