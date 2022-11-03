package main

import (
	balances_http "crypto-calc/cmd/http/balance/get_balances_sum"
	useCases "crypto-calc/core/use-cases"
	"crypto-calc/packages/etherscan"
	"crypto-calc/packages/infura"
	"flag"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvOrThrow(key string, providerName string) string {
	err := godotenv.Load(providerName + ".env")

	if err != nil {
		panic("Error loading .env file")
	}

	return os.Getenv(key)
}

var rateLimit int
var numRequestsInParallel int
var providerName string

func init() {
	flag.IntVar(&rateLimit, "rate-limit", 1, "Number of requests in parallel")
	flag.IntVar(&numRequestsInParallel, "parallelism", 2, "Number of requests in time interval")
	flag.StringVar(&providerName, "providerName", "infura", "Number of requests in time interval")
}

func main() {
	flag.Parse()

	var apiKey = GetEnvOrThrow("API_KEY", providerName)
	var providersMap = map[string]useCases.BalanceProvider{"infura": infura.Infura{ApiKey: apiKey}, "etherscan": etherscan.Etherscan{ApiKey: apiKey}}

	balanceProvider := providersMap[providerName]

	http.HandleFunc("/balance", balances_http.GetBalancesSumHandler(balanceProvider, numRequestsInParallel, rateLimit))

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		panic(err)
	}
}
