package main

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type BalanceProviderMock struct {
	mock.Mock
}

func (b BalanceProviderMock) GetEtherBalance() (int, error) {
	args := b.Called()
	balance := 44
	return balance, args.Error(1)
}

func TestCli(t *testing.T) {
	var addresses = []string{"0xab5801a7d398351b8be11c439e05c5b3259aec9b", "0xCb235D0dc69E8D085b4179c77E7981D1B9D90ACA"}

	balanceProvider := BalanceProviderMock{}
	balanceProvider.On("GetEtherBalance").Return(10, nil)

	GetBalancesSum(addresses, balanceProvider)

	balanceProvider.AssertCalled(t, "GetEtherBalance", "0xab5801a7d398351b8be11c439e05c5b3259aec9b")
	balanceProvider.AssertCalled(t, "GetEtherBalance", "0xCb235D0dc69E8D085b4179c77E7981D1B9D90ACA")
}
