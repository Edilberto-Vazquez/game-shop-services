package apigateway

import (
	"context"
	"log"
	"os"

	"github.com/Edilberto-Vazquez/game-shop-services/apigateway/drivers/http/routes"
	"github.com/Edilberto-Vazquez/game-shop-services/apigateway/drivers/http/server"
	"github.com/joho/godotenv"
)

func StartApiGateWay() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port: PORT,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(routes.GetRoutes)
}
