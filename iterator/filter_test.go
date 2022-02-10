package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIteration(t *testing.T) {
	t.Run("Original iterable not modified", func(t *testing.T) {
		slice := []int{-1, 2, 3}

		sliceIter := NewSliceIterator(slice)
		filterIterator := NewFilterIterator(sliceIter, func(t int) bool { return t > 0 })

		filterIterator.MoveNext()
		assert.Equal(t, -1, filterIterator.GetCurrent())

		filterIterator.MoveNext()
		assert.Equal(t, 2, filterIterator.GetCurrent())

		filterIterator.MoveNext()
		assert.Equal(t, 3, filterIterator.GetCurrent())
	})

	t.Run("Empty slice - predicate: > 0", func(t *testing.T) {
		slice := []int{}

		sliceIter := NewSliceIterator(slice)
		filterIterator := NewFilterIterator(sliceIter, func(t int) bool { return t > 0 })

		assert.False(t, filterIterator.MoveNext())
	})

	t.Run("[1,2,3] - predicate: > 0", func(t *testing.T) {
		slice := []int{1, 2, 3}

		sliceIter := NewSliceIterator(slice)
		filterIterator := NewFilterIterator(sliceIter, func(t int) bool { return t > 0 })

		filterIterator.MoveNext()
		assert.Equal(t, 1, filterIterator.GetCurrent())

		filterIterator.MoveNext()
		assert.Equal(t, 2, filterIterator.GetCurrent())

		filterIterator.MoveNext()
		assert.Equal(t, 3, filterIterator.GetCurrent())
	})

	t.Run("[1,2,3] - predicate: > 1", func(t *testing.T) {
		slice := []int{1, 2, 3}

		sliceIter := NewSliceIterator(slice)
		filterIterator := NewFilterIterator(sliceIter, func(t int) bool { return t > 1 })

		filterIterator.MoveNext()
		assert.Equal(t, 2, filterIterator.GetCurrent())

		filterIterator.MoveNext()
		assert.Equal(t, 3, filterIterator.GetCurrent())
	})
}

func BenchmarkFilter(b *testing.B) {
	b.Run("Create Filter iterable", func(b *testing.B) {
		slice := []int{1, 2, 3}

		sliceIterator := NewSliceIterator(slice)

		for i := b.N; i < b.N; i++ {
			NewFilterIterator(sliceIterator, func(t int) bool { return t > 0 })
		}
	})

	b.Run("Iterate over 10 element iterable", func(b *testing.B) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		sliceIter := NewSliceIterator(slice)
		filterIterator := NewFilterIterator(sliceIter, func(t int) bool { return t > 0 })

		for i := b.N; i < b.N; i++ {
			for filterIterator.MoveNext() {
				filterIterator.GetCurrent()
			}
		}
	})
}
