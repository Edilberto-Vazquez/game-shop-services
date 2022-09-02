package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sessionRoutes(rg *gin.RouterGroup) {
	session := rg.Group("/session")

	session.POST("/signup", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "signup")
	})

	session.GET("/login", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "login")
	})
}
