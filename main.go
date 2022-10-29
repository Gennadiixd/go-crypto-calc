package main

import (
	"log"
)

var apiKey = config.API_KEY
var providersMap = map[string]BalanceProvider{"infura": Infura{apiKey: apiKey}, "etherscan": Etherscan{apiKey: apiKey}}
var addresses = []string{"0xab5801a7d398351b8be11c439e05c5b3259aec9b", "0xCb235D0dc69E8D085b4179c77E7981D1B9D90ACA"}

func main() {
	balanceProvider := providersMap[providerName]
	var balance, err = getBalancesSum(addresses, balanceProvider)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(balance)
}
