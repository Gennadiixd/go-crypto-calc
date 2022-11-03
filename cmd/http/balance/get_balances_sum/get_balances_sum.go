package balances_http

import (
	useCases "crypto-calc/core/use-cases"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func GetBalancesSumHandler(balanceProvider useCases.BalanceProvider, numRequestsInParallel int, rateLimit int) func(w http.ResponseWriter, r *http.Request) {
	balanceHandler := func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		addresses := r.Form["addresses"]

		var balance, err = useCases.GetBalancesSum(addresses, balanceProvider, numRequestsInParallel, rateLimit)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("got /balance request\n")

		var balanceStr = strconv.FormatFloat(balance, 'f', 6, 64)
		if err != nil {
			panic(err)
		}

		io.WriteString(w, balanceStr)
	}

	return balanceHandler
}
