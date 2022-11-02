package etherscan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
)

type Etherscan struct {
	ApiKey string
}

type etherscan_resp struct {
	Result  string `json:"result"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func (e Etherscan) GetEtherBalance(address string) (float64, error) {
	requestURL := fmt.Sprintf("https://api.etherscan.io/api?module=account&action=balance&address=%s&tag=latest&apikey=%s", address, e.ApiKey)

	resp, err := http.Get(requestURL)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var respData etherscan_resp

	err = json.Unmarshal(body, &respData)
	if err != nil {
		panic(err)
	}

	balance, err := strconv.ParseFloat(respData.Result, 64)
	if err != nil {
		panic(err)
	}

	return balance / math.Pow(10, 18), nil
}
