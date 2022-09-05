package middleware

import (
	"net/http"

	"github.com/Edilberto-Vazquez/game-shop-services/src/apigateway/drivers/http/server"
	"github.com/Edilberto-Vazquez/game-shop-services/src/apigateway/utils"
	"github.com/gin-gonic/gin"
)

var (
	NO_AUTH_NEEDED = []string{
		"login",
		"signup",
	}
)

func CheckAuthMiddleware(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !utils.RouteNeedToken(ctx.FullPath(), NO_AUTH_NEEDED) {
			ctx.Next()
			return
		}
		token, err := utils.ProcessToken(ctx.GetHeader("Authorization"), s)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		ctx.Set("token", token)
		ctx.Next()
	}
}
