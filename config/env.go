package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func EnvMongoUri() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	return os.Getenv("MONGO_URI")
}