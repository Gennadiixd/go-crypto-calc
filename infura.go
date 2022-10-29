package main

type Infura struct {
	apiKey string
}

func (i Infura) getEtherBalance(address string) (int, error) {
	return 11, nil
}
