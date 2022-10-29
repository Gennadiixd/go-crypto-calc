package infura

type Infura struct {
	ApiKey string
}

func (i Infura) GetEtherBalance(address string) (int, error) {
	return 11, nil
}
