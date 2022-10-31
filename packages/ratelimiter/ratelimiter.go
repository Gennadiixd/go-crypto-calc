package ratelimiter

import (
	"fmt"
	"time"
)

func Ratelimiter[T, U any](args []T, fn func(T) (U, error), numRequestsInParallel int, rateLimit int) ([]U, error) {

	var balances []U

	fnsChan := make(chan func() (U, error), numRequestsInParallel)
	resultsChan := make(chan U)

	for i := 0; i < cap(fnsChan); i++ {
		go worker(fnsChan, resultsChan, time.Duration(rateLimit))
	}

	go func() {

		i := 0
		for i <= 3 {
			arg := args[i]
			f := func() (U, error) {
				fmt.Printf("Request start ")
				fmt.Println(arg)
				return fn(arg)
			}

			fnsChan <- f
			i = i + 1
		}

		// for i := 0; i < len(args); i++ {

		// 	// fmt.Println(i)
		// 	f := func() (U, error) {
		// 		fmt.Println(i)
		// 		return fn(args[i])
		// 	}

		// 	fnsChan <- f
		// }
	}()

	for i := 0; i < len(args); i++ {
		balance := <-resultsChan
		balances = append(balances, balance)
	}

	close(fnsChan)
	close(resultsChan)

	return balances, nil
}

func worker[U any](fnsChan <-chan func() (U, error), resultsChan chan<- U, rateLimit time.Duration) {
	for fn := range fnsChan {
		timer := time.NewTimer(rateLimit * time.Second)

		result, err := fn()
		if err != nil {
			panic(err)
		}

		<-timer.C
		resultsChan <- result
		fmt.Println("Request end")
	}
}
