package routes

import (
	"net/http"

	sess "github.com/Edilberto-Vazquez/game-shop-services/session"
	"github.com/Edilberto-Vazquez/game-shop-services/session/models"
	"github.com/gin-gonic/gin"
)

func sessionRoutes(rg *gin.RouterGroup) {
	session := rg.Group("/session")

	session.POST("/signup", func(ctx *gin.Context) {
		code := sess.NewSessionService()
		code.SignUp(ctx, &models.User{
			Id:        "asdqwf2423r231",
			UserName:  "Edi",
			Email:     "edi@mail.com",
			CountryId: "MX",
			Salt:      "dy2j1",
			Hash:      "dh7828fidw8872g423b4",
		})
		ctx.JSON(http.StatusOK, "signup")
	})

	session.GET("/login", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "login")
	})
}
