package core

import (
	"crypto-calc/packages/ratelimiter"
)

type BalanceProvider interface {
	GetEtherBalance(string) (float64, error)
}

func GetBalancesSum(addresses []string, balanceProvider BalanceProvider, numRequestsInParallel int, rateLimit int) (float64, error) {
	var getEtherBalanceLimited = ratelimiter.TimeLimiter(balanceProvider.GetEtherBalance, rateLimit)
	var balances, err = ratelimiter.Ratelimiter(addresses, getEtherBalanceLimited, numRequestsInParallel, rateLimit)
	if err != nil {
		panic(err)
	}

	var result float64

	for _, balance := range balances {
		result = result + balance
	}

	return result, nil
}
