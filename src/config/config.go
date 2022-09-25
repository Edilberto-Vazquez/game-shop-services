package config

import (
	"os"
)

var (
	Env string
)

func SetEnvironment() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal(err)
	// }
	switch os.Getenv("GIN_MODE") {
	case "release":
		Env = ""
	case "debug":
		Env = "DB_"
	case "test":
		Env = "TEST_"
	default:
		Env = "DEV_"
	}
}
