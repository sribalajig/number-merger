package numbers

import (
	"testing"
	"time"
)

/*All of these tests run in parallel*/

func Test_ForEmptyResult(t *testing.T) {
	t.Parallel()

	timer := time.NewTimer(time.Millisecond * 500)

	mergedArray := NewAccumulator(make(chan *[]int), 4).Accumulate(timer)

	if mergedArray == nil {
		t.Error("Expected array, got nil")
		t.Fail()
	}

	if len(*mergedArray) != 0 {
		t.Errorf("Expected empty array, got array of length %d", len(*mergedArray))
	}
}

func Test_SingleResultReturnsAfter500ms(t *testing.T) {
	t.Parallel()

	results := make(chan *[]int)

	timer := time.NewTimer(time.Millisecond * 500)

	// this should time out
	go func() {
		time.Sleep(time.Millisecond * 510)
		results <- &[]int{1, 2, 4}
	}()

	mergedArray := NewAccumulator(results, 4).Accumulate(timer)

	if mergedArray == nil {
		t.Error("Expected array, got nil")
		t.Fail()
	}

	if len(*mergedArray) != 0 {
		t.Errorf("Expected empty array, got array of length %d", len(*mergedArray))
	}
}

func Test_OneResultReturnsWithin500ms(t *testing.T) {
	t.Parallel()

	results := make(chan *[]int)
	timer := time.NewTimer(time.Millisecond * 500)

	// this should time out
	go func() {
		time.Sleep(time.Millisecond * 510)
		results <- &[]int{1, 2, 4}
	}()

	// this should not time out
	go func() {
		time.Sleep(time.Millisecond * 490)
		results <- &[]int{4, 5, 6}
	}()

	mergedArray := NewAccumulator(results, 4).Accumulate(timer)

	if mergedArray == nil {
		t.Error("Expected array, got nil")
		t.Fail()
	}

	if len(*mergedArray) != 3 {
		t.Errorf("Expected array of length %d, got array of length %d", 3, len(*mergedArray))
		t.Fail()
	}
}

func Test_AllResultsReturnIn500ms(t *testing.T) {
	t.Parallel()

	results := make(chan *[]int)

	// this should not time out
	go func() {
		time.Sleep(time.Millisecond * 490)
		results <- &[]int{1, 2, 3}
	}()

	// this should not time out
	go func() {
		time.Sleep(time.Millisecond * 120)
		results <- &[]int{4, 5, 6}
	}()

	// this should not time out
	go func() {
		time.Sleep(time.Millisecond * 300)
		results <- &[]int{7, 8, 9}
	}()

	// this should not time out
	go func() {
		time.Sleep(time.Millisecond * 495)
		results <- &[]int{10, 11, 12}
	}()

	timer := time.NewTimer(time.Millisecond * 500)

	mergedArray := NewAccumulator(results, 4).Accumulate(timer)

	if len(*mergedArray) != 12 {
		t.Errorf("Expected array of length %d, got array of length %d", 12, len(*mergedArray))
		t.Fail()
	}
}
