package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadDbUrl() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load .env file", err)
	} else {
		return os.Getenv("MONGODB_URL")
	}
	return ""
}
