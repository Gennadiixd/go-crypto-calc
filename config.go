package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var providerName = getCliArgOrThrow("provider")

func getEnvOrThrow(key string) string {
	err := godotenv.Load(providerName + ".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func getCliArgOrThrow(argName string) string {
	for i := range os.Args {
		var arg = os.Args[i]

		if strings.Index(arg, "--"+argName) == 0 {
			return strings.Split(arg, "=")[1]
		}
	}

	log.Fatalf("Error finding cli argument" + argName)

	return ""
}

type Config struct {
	API_KEY     string
	parallelism int
	rateLimit   int
}

var rateLimit, _ = strconv.Atoi(getCliArgOrThrow("rate-limit"))
var parallelism, _ = strconv.Atoi(getCliArgOrThrow("parallelism"))

var config = &Config{
	API_KEY:     getEnvOrThrow("API_KEY"),
	parallelism: parallelism,
	rateLimit:   rateLimit,
}
