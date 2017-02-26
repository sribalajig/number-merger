package numbers

import (
	"log"
	"time"
)

/*Accumulator accumulates results from the various HTTP requests*/
type Accumulator struct {
	results      <-chan *[]int
	resultLength int
}

func NewAccumulator(results <-chan *[]int, resultLength int) Accumulator {
	return Accumulator{
		results:      results,
		resultLength: resultLength,
	}
}

/*Accumulate listens on the results channel. When a result is available, it will try and
merge the result within the timout period. It returns an array of integers which
represents the final merged array.*/
func (accumulator Accumulator) Accumulate(timer *time.Timer) *[]int {
	mergedArray := []int{}

	for i := 0; i < accumulator.resultLength; i++ {
		select {
		case <-timer.C:
			log.Printf("Time out after processing %d arrays, exiting\n", i)

			return &mergedArray

		case result := <-accumulator.results:
			mergedResult := make(chan MergedResult)

			go MergeWithinTimeLimit(mergedArray, *result, mergedResult, timer)

			res := <-mergedResult

			if res.timedOut {
				return &res.result
			}

			mergedArray = res.result
		}
	}

	log.Println("Processed all available arrays before timeout")

	return &mergedArray
}
