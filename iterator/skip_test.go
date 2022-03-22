package iterator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSKipIteration(t *testing.T) {
	testCases := []struct {
		slice    []int
		skip     int
		expected []int
	}{
		{
			slice:    []int{},
			skip:     0,
			expected: []int{},
		},
		{
			slice:    []int{},
			skip:     1,
			expected: []int{},
		},
		{
			slice:    []int{1, 2, 3},
			skip:     0,
			expected: []int{1, 2, 3},
		},
		{
			slice:    []int{1, 2, 3},
			skip:     1,
			expected: []int{2, 3},
		},
		{
			slice:    []int{1, 2, 3},
			skip:     2,
			expected: []int{3},
		},
		{
			slice:    []int{1, 2, 3},
			skip:     3,
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("slice: %v, skip: %v, actual: %v", tc.slice, tc.skip, tc.expected), func(t *testing.T) {
			sliceIterator := NewSliceIterator(tc.slice)
			takeIterator := NewSkipIterator(sliceIterator, tc.skip)

			actual := iterateToEnd(takeIterator)
			assert.EqualValues(t, tc.expected, actual)
		})
	}
}

func iterateToEnd[T any](iter Iterator[T]) []T {
	result := []T{}
	for iter.MoveNext() {
		result = append(result, iter.GetCurrent())
	}

	return result
}
