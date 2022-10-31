package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var providerName = os.Args[3]

func getEnvOrThrow(key string) string {
	err := godotenv.Load(providerName + ".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

type Config struct {
	API_KEY string
}

var config = &Config{
	API_KEY: getEnvOrThrow("API_KEY"),
}
