package main

type Etherscan struct {
	apiKey string
}

func (e Etherscan) getEtherBalance(address string) (int, error) {
	return 90, nil
}
