package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func SetEnvironment() (environment string) {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	switch os.Getenv("APP_ENV") {
	case "production":
		environment = ""
	default:
		environment = "DEV_"
	}
	return
}
