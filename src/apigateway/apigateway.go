package apigateway

import (
	"context"
	"log"
	"os"

	"github.com/Edilberto-Vazquez/game-shop-services/src/apigateway/config"
	"github.com/Edilberto-Vazquez/game-shop-services/src/apigateway/drivers/http/routes"
	"github.com/Edilberto-Vazquez/game-shop-services/src/apigateway/drivers/http/server"
)

func StartApiGateWay() {
	prefix := config.SetEnvironment()

	PORT := os.Getenv(prefix + "PORT")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port: PORT,
	})

	if err != nil {
		log.Fatal(err)
	}
	s.Start(routes.GetRoutes)
}