package core

import (
	"crypto-calc/packages/batcher"
)

type BalanceProvider interface {
	GetEtherBalance(string) (float64, error)
}

func GetBalancesSum(addresses []string, balanceProvider BalanceProvider) (float64, error) {
	var balances, err = batcher.Batcher(addresses, balanceProvider.GetEtherBalance, 2)
	if err != nil {
		panic(err)
	}

	var result float64

	for _, balance := range balances {
		result = result + balance
	}

	return result, nil
}
