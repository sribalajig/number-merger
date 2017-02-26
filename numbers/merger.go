package numbers

import (
	"time"
)

/*MergeWithinTimeLimit takes two arrays and publishes the result into
a channel within the specified time limit. If time limit expires,
the original array is pushed into the channel*/
func MergeWithinTimeLimit(
	originalArray []int,
	arrayToMerge []int,
	result chan<- MergedResult,
	timer *time.Timer) {
	mergedResult := make(chan []int)

	go Merge(originalArray, arrayToMerge, mergedResult)

	select {
	case <-timer.C:
		result <- MergedResult{
			result:   originalArray,
			timedOut: true,
		}
	case r := <-mergedResult:
		result <- MergedResult{
			result:   r,
			timedOut: false,
		}
	}
}
