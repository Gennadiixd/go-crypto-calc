package core

type BalanceProvider interface {
	GetEtherBalance(string) (int, error)
}

func GetBalancesSum(addresses []string, balanceProvider BalanceProvider) (int, error) {
	var balanceTtl = 0

	for _, address := range addresses {
		var balance, err = balanceProvider.GetEtherBalance(address)

		if err != nil {
			return 0, err
		}

		balanceTtl = balanceTtl + balance
	}

	return balanceTtl, nil
}
