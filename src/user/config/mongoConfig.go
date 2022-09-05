package config

import (
	"os"

	"github.com/Edilberto-Vazquez/game-shop-services/src/user/constants"
)

type DBConfig struct {
	Name       string
	Collection string
	URI        string
}

func NewMongoConfig() DBConfig {
	prefix := SetEnvironment()

	uri := os.Getenv(prefix + "MONGODB_URI")

	return DBConfig{
		Name:       constants.DB_NAME,
		Collection: constants.COLLECTION_NAME,
		URI:        uri,
	}
}
