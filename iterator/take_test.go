package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrent(t *testing.T) {
	t.Run("Empty slice - take 1", func(t *testing.T) {
		slice := []int{}
		sliceIterator := NewSliceIterator(slice)
		takeIterator := NewTakeIterator(sliceIterator, 1)

		assert.False(t, takeIterator.MoveNext())
	})

	t.Run("[1,2,3] - take 1", func(t *testing.T) {
		slice := []int{1, 2, 3}
		sliceIterator := NewSliceIterator(slice)
		takeIterator := NewTakeIterator(sliceIterator, 1)

		assert.True(t, takeIterator.MoveNext())
		assert.Equal(t, 1, takeIterator.GetCurrent())

		assert.False(t, takeIterator.MoveNext())
	})

	t.Run("[1,2,3] - take 2", func(t *testing.T) {
		slice := []int{1, 2, 3}
		sliceIterator := NewSliceIterator(slice)
		takeIterator := NewTakeIterator(sliceIterator, 2)

		assert.True(t, takeIterator.MoveNext())
		assert.Equal(t, 1, takeIterator.GetCurrent())

		assert.True(t, takeIterator.MoveNext())
		assert.Equal(t, 2, takeIterator.GetCurrent())

		assert.False(t, takeIterator.MoveNext())
	})
}

func TestNewTakeIterator(t *testing.T) {
	t.Run("Nil inner iterator", func(t *testing.T) {
		innerIter := Iterator[int](nil)
		take := 2

		act := func() {
			NewTakeIterator(innerIter, take)
		}

		assert.Panics(t, act, "Expected 'innerIter' to be Iterator but received 'nil'")
	})

	t.Run("Negative take parameter", func(t *testing.T) {
		slice := []int{1, 2, 3}
		iter := NewSliceIterator(slice)
		take := -1

		act := func() {
			NewTakeIterator(iter, take)
		}

		assert.Panics(t, act, "Expected 'slice' to be non-negative number but received: -1")
	})
}
