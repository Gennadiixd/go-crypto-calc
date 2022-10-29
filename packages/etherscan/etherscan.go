package etherscan

type Etherscan struct {
	ApiKey string
}

func (e Etherscan) GetEtherBalance(address string) (int, error) {
	return 90, nil
}
