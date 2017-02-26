package numbers

import (
	"testing"
)

var testCases = []struct {
	a        []int
	b        []int
	expected []int
}{
	{[]int{}, []int{}, []int{}},                                       // two empty arrays
	{[]int{1, 2, 3}, []int{1, 2}, []int{1, 2, 3}},                     // there are two intersecting elements
	{[]int{1, 1, 1}, []int{1}, []int{1}},                              // only one element is expected in the intersecting array
	{[]int{1, 1, 2, 2, 3, 3}, []int{}, []int{1, 2, 3}},                // one empty array
	{[]int{}, []int{1, 1, 2, 2, 3, 3, 3, 4, 4, 4}, []int{1, 2, 3, 4}}, // one empty array
	{[]int{1, 1, 2, 2, 3, 3}, []int{1, 2, 3}, []int{1, 2, 3}},         // both arrays have intersecting elements
	{[]int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},         // no intersecting elements between the arrays
}

func TestMerge(t *testing.T) {
	for _, testCase := range testCases {
		merged := make(chan []int)

		go Merge(testCase.a, testCase.b, merged)

		result := <-merged

		if len(result) != len(testCase.expected) {
			t.Errorf("Expected merged array length %d, got %d", len(testCase.expected), len(result))
			t.Errorf("Expected : %q", testCase.expected)
			t.Errorf("Got : %q", result)
			t.Fail()

			continue
		}

		for index, actual := range result {
			if result[index] != actual {
				t.Errorf("Expected %d at index %d, got %d", result[index], index, actual)

				t.Fail()

				break
			}
		}
	}
}
