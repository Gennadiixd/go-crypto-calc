package batcher

import (
	"math"
	"sync"
)

func Batcher[T, U any](args []T, fn func(T) (U, error), maxBatchSize int) ([]U, error) {
	results := make([]U, len(args))

	skip := 0
	argsAmount := len(args)
	batchAmount := int(math.Ceil(float64(argsAmount / maxBatchSize)))

	for i := 0; i <= batchAmount; i++ {
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

				results[idx] = result
			}(batchItems[idx], idx)
		}

		itemProcessingGroup.Wait()
	}

	return results, nil
}
