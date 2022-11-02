package ratelimiter

import (
	"fmt"
	"time"
)

func Ratelimiter[T, U any](args []T, fn func(T) (U, error), numRequestsInParallel int, rateLimit int) ([]U, error) {

	var results []U

	fnsChan := make(chan func() (U, error), numRequestsInParallel)
	resultsChan := make(chan U)

	for i := 0; i < cap(fnsChan); i++ {
		go worker(fnsChan, resultsChan, time.Duration(rateLimit))
	}

	go func() {

		i := 0
		for i < len(args) {
			arg := args[i]
			f := func() (U, error) {
				fmt.Printf("Request start ")
				fmt.Println(arg)
				return fn(arg)
			}

			fnsChan <- f
			i = i + 1
		}
	}()

	for i := 0; i < len(args); i++ {
		balance := <-resultsChan
		results = append(results, balance)
	}

	close(fnsChan)
	close(resultsChan)

	return results, nil
}

func TimeLimiter[T, U any](fn func(T) (U, error), rateLimit int) func(T) (U, error) {

	limited := func(arg T) (U, error) {
		timer := time.NewTimer(time.Duration(rateLimit) * time.Second)

		result, err := fn(arg)
		if err != nil {
			panic(err)
		}

		<-timer.C
		return result, nil
	}

	return limited
}

func worker[U any](fnsChan <-chan func() (U, error), resultsChan chan<- U, rateLimit time.Duration) {
	for fn := range fnsChan {

		result, err := fn()
		if err != nil {
			panic(err)
		}

		resultsChan <- result
		fmt.Println("Request end")
	}
}
