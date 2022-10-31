package main

import (
	useCases "crypto-calc/core/use-cases"
	"crypto-calc/packages/etherscan"
	"crypto-calc/packages/infura"
	"flag"
	"log"
)

var apiKey = config.API_KEY
var providersMap = map[string]useCases.BalanceProvider{"infura": infura.Infura{ApiKey: apiKey}, "etherscan": etherscan.Etherscan{ApiKey: apiKey}}
var addresses = []string{"0xab5801a7d398351b8be11c439e05c5b3259aec9b", "0xCb235D0dc69E8D085b4179c77E7981D1B9D90ACA", "0xb120c885f1527394C78D50e7C7DA57DEfb24F612", "0x297BF847Dcb01f3e870515628b36EAbad491e5E8"}

var rateLimit int
var numRequestsInParallel int

func init() {
	flag.IntVar(&rateLimit, "rate-limit", 1, "Number of requests in parallel")
	flag.IntVar(&numRequestsInParallel, "parallelism", 2, "Number of requests in time interval")
}

func main() {
	flag.Parse()

	balanceProvider := providersMap[providerName]
	var balance, err = useCases.GetBalancesSum(addresses, balanceProvider, numRequestsInParallel, rateLimit)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(balance)
}
