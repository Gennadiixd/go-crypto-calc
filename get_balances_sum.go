package main

type BalanceProvider interface {
	getEtherBalance(string) (int, error)
}

func getBalancesSum(addresses []string, balanceProvider BalanceProvider) (int, error) {
	var balanceTtl = 0

	for _, address := range addresses {
		var balance, err = balanceProvider.getEtherBalance(address)

		if err != nil {
			return 0, err
		}

		balanceTtl = balanceTtl + balance
	}

	return balanceTtl, nil
}
