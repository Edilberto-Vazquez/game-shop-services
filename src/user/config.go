package user

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	dbName       string
	dbCollection string
	dbUri        string
}

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

func NewConfig() (Config, error) {
	prefix := SetEnvironment()
	uri := os.Getenv(prefix + "MONGODB_URI")
	return Config{
		dbName:       DB_NAME,
		dbCollection: DB_COLLECTION,
		dbUri:        uri,
	}, nil
}
