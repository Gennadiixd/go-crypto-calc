package batcher

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func Batcher[T, U any](args []T, fn func(T) (U, error), maxBatchSize int) ([]U, error) {
	results := make([]U, len(args))

	skip := 0
	argsAmount := len(args)
	batchAmount := int(math.Ceil(float64(argsAmount / maxBatchSize)))

	for i := 0; i < batchAmount; i++ {
		timer := time.NewTimer(3 * time.Second)

		lowerBound := skip
		upperBound := skip + maxBatchSize

		if upperBound > argsAmount {
			upperBound = argsAmount
		}

		batchItems := args[lowerBound:upperBound]

		skip += maxBatchSize

		var itemProcessingGroup sync.WaitGroup
		itemProcessingGroup.Add(len(batchItems))

		for idx := range batchItems {
			go func(currentArg T, idx int) {
				defer itemProcessingGroup.Done()

				var result, err = fn(currentArg)
				if err != nil {
					panic(err)
				}
				fmt.Println(result)
				results[idx+i] = result

			}(batchItems[idx], idx)
		}

		<-timer.C

		fmt.Println("Batch done")
		itemProcessingGroup.Wait()
	}

	return results, nil
}
