package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load() // Load secreats from .env
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
