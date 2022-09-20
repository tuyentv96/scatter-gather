package scattergather

import (
	"errors"
	"math"
	"sync"

	"golang.org/x/sync/errgroup"
)

var (
	ErrInvalidBatchSize = errors.New("batch size must be great than zero")
)

// ScattergatherWithFuncs executes multiple futures funtions
// Return error if anfuture has failed.
func ScattergatherWithFuncs[T any](futures [](func() ([]T, error))) ([]T, error) {
	g := &errgroup.Group{}
	mu := &sync.Mutex{}
	var result []T

	for _, f := range futures {
		f := f // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			rs, err := f()
			if err != nil {
				return err
			}

			if len(rs) > 0 {
				mu.Lock()
				result = append(result, rs...)
				mu.Unlock()
			}

			return nil
		})
	}

	// Wait for all futures are completed.
	if err := g.Wait(); err != nil {
		return nil, err
	}

	return result, nil
}

// ScattergatherWithInputParams executes with input params array.
// Input params array will be partitioned by batchSize
// Each batch is executed by fn
// Return errors if any params is failed.
func ScattergatherWithInputParams[K, V any](params []K, batchSize int, fn func([]K) ([]V, error)) ([]V, error) {
	if batchSize <= 0 {
		return nil, ErrInvalidBatchSize
	}

	var futures [](func() ([]V, error))

	// Build futures array
	// Partition params into batches by batchSize
	for i := 0; i < len(params); i += batchSize {
		batch := params[i:int(math.Min(float64(i+batchSize), float64(len(params))))]
		futures = append(futures, func() ([]V, error) {
			return fn(batch)
		})
	}

	return ScattergatherWithFuncs(futures)
}
