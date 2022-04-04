package iterator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkipIteration(t *testing.T) {
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
			skipIterator := NewSkipIterator(sliceIterator, tc.skip)

			actual := iterateToEnd(skipIterator)
			assert.EqualValues(t, tc.expected, actual)
		})
	}
}

func TestNewSkipIterator(t *testing.T) {
	t.Run("Nil inner iterator", func(t *testing.T) {
		sliceIter := Iterator[int](nil)
		skip := 1

		act := func() {
			NewSkipIterator(sliceIter, skip)
		}

		assert.Panics(t, act, "Expected 'innerIter' to be Iterator but received 'nil'")

	})

	t.Run("Negative skip parameter", func(t *testing.T) {
		slice := []int{1, 2, 3}
		sliceIterator := NewSliceIterator(slice)

		act := func() {
			NewSkipIterator(sliceIterator, -1)
		}

		assert.Panics(t, act, "Expected 'slice' to be non-negative number but received: -1")
	})
}

func iterateToEnd[T any](iter Iterator[T]) []T {
	result := []T{}
	for iter.MoveNext() {
		result = append(result, iter.GetCurrent())
	}

	return result
}
