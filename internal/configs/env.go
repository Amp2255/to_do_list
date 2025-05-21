package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadDbUrl() string {
	var err error
	log.Printf("os.Getenv(ENV) :: %s", os.Getenv("ENV"))
	if os.Getenv("ENV") != "production" {
		err = godotenv.Load()
	}
	//err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load .env file", err)
	} else {
		return os.Getenv("MONGODB_URL")
	}
	return ""
}

func LoadPort() string {
	var err error
	log.Printf("os.Getenv(ENV) :: %s", os.Getenv("ENV"))
	if os.Getenv("ENV") != "production" {
		err = godotenv.Load()
	}
	if err != nil {
		log.Fatal("Failed to load .env file", err)
	} else {
		return os.Getenv("PORT")
	}
	return "8080"
}

func LoadWebUrl() string {
	var err error
	if os.Getenv("ENV") != "production" {
		err = godotenv.Load()
	}
	if err != nil {
		log.Fatal("Failed to load .env file", err)
	} else {
		return os.Getenv("WEBURL")
	}
	return "0.0.0.0"
}
