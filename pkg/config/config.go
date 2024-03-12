package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	env := os.Getenv("ENV")
	if err != nil && env == "" {
		log.Println("failed to load env variables")
	}
}
