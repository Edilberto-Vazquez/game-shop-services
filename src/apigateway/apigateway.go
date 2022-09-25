package apigateway

import (
	"context"
	"log"
	"os"

	"github.com/Edilberto-Vazquez/game-shop-services/src/apigateway/drivers/http/routes"
	"github.com/Edilberto-Vazquez/game-shop-services/src/apigateway/drivers/http/server"
	"github.com/Edilberto-Vazquez/game-shop-services/src/apigateway/drivers/http/services"
	"github.com/Edilberto-Vazquez/game-shop-services/src/config"
)

func StartApiGateWay() {
	config.SetEnvironment()
	log.Println(config.Env + "PORT")
	PORT := os.Getenv(config.Env + "PORT")
	s, err := server.NewServer(
		context.Background(),
		&server.Config{Port: PORT},
		services.NewServices(),
	)
	if err != nil {
		log.Fatal(err)
	}
	s.Start(routes.GetRoutes)
}
