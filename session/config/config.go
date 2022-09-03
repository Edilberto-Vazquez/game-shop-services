package config

import (
	"log"
	"os"

	"github.com/Edilberto-Vazquez/game-shop-services/session/constants"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	Name       string
	Collection string
	URI        string
}

func Config() DBConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	uri := os.Getenv("MONGODB_URI")

	return DBConfig{
		Name:       constants.DB_NAME,
		Collection: constants.COLLECTION_NAME,
		URI:        uri,
	}
}
