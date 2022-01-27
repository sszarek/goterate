package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIteration(t *testing.T) {
	t.Run("Returns new iterable", func(t *testing.T) {
		slice := []int{1, 2, 3}

		iterable := NewSliceIterable(slice)
		filteredIterable := iterable.Where(func(t int) bool { return t > 0 })

		assert.NotSame(t, iterable, filteredIterable)
	})

	t.Run("Original iterable not modified", func(t *testing.T) {
		slice := []int{-1, 2, 3}

		iterable := NewSliceIterable(slice)
		iterable.Where(func(t int) bool { return t > 0 })

		iterator := iterable.GetIterator()

		iterator.MoveNext()
		assert.Equal(t, -1, iterator.GetCurrent())

		iterator.MoveNext()
		assert.Equal(t, 2, iterator.GetCurrent())

		iterator.MoveNext()
		assert.Equal(t, 3, iterator.GetCurrent())
	})

	t.Run("Empty slice - predicate: > 0", func(t *testing.T) {
		slice := []int{}

		iterable := NewSliceIterable(slice).Where(func(t int) bool { return t > 0 })
		iterator := iterable.GetIterator()

		assert.False(t, iterator.MoveNext())
	})

	t.Run("[1,2,3] - predicate: > 0", func(t *testing.T) {
		slice := []int{1, 2, 3}

		iterable := NewSliceIterable(slice).Where(func(t int) bool { return t > 0 })
		iterator := iterable.GetIterator()

		iterator.MoveNext()
		assert.Equal(t, 1, iterator.GetCurrent())

		iterator.MoveNext()
		assert.Equal(t, 2, iterator.GetCurrent())

		iterator.MoveNext()
		assert.Equal(t, 3, iterator.GetCurrent())
	})

	t.Run("[1,2,3] - predicate: > 1", func(t *testing.T) {
		slice := []int{1, 2, 3}

		iterable := NewSliceIterable(slice).Where(func(t int) bool { return t > 1 })
		iterator := iterable.GetIterator()

		iterator.MoveNext()
		assert.Equal(t, 2, iterator.GetCurrent())

		iterator.MoveNext()
		assert.Equal(t, 3, iterator.GetCurrent())
	})
}

func BenchmarkWhere(b *testing.B) {
	b.Run("Create Where iterable", func(b *testing.B) {
		slice := []int{1, 2, 3}

		iterable := NewSliceIterable(slice)

		for i := b.N; i < b.N; i++ {
			iterable.Where(func(t int) bool { return t > 0 })
		}
	})

	b.Run("Iterate over 10 element iterable", func(b *testing.B) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		iterable := NewSliceIterable(slice).Where(func(t int) bool { return t > 0 })
		iterator := iterable.GetIterator()

		for i := b.N; i < b.N; i++ {
			for iterator.MoveNext() {
				iterator.GetCurrent()
			}
		}
	})
}
