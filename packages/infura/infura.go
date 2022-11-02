package infura

type Infura struct {
	ApiKey string
}

func (i Infura) GetEtherBalance(address string) (float64, error) {
	return 11, nil
}
