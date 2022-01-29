package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrent(t *testing.T) {
	t.Run("Returns new iterable", func(t *testing.T) {
		slice := []int{1, 2, 3}
		iterable := NewSliceIterable(slice)
		takeIterable := iterable.Take(1)

		assert.NotSame(t, iterable, takeIterable)
	})

	t.Run("Empty slice - take 1", func(t *testing.T) {
		slice := []int{}
		iterable := NewSliceIterable(slice).Take(1)
		iterator := iterable.GetIterator()

		assert.False(t, iterator.MoveNext())
	})

	t.Run("[1,2,3] - take 1", func(t *testing.T) {
		slice := []int{1, 2, 3}
		iterable := NewSliceIterable(slice).Take(1)
		iterator := iterable.GetIterator()

		assert.True(t, iterator.MoveNext())
		assert.Equal(t, 1, iterator.GetCurrent())

		assert.False(t, iterator.MoveNext())
	})

	t.Run("[1,2,3] - take 2", func(t *testing.T) {
		slice := []int{1, 2, 3}
		iterable := NewSliceIterable(slice).Take(2)
		iterator := iterable.GetIterator()

		assert.True(t, iterator.MoveNext())
		assert.Equal(t, 1, iterator.GetCurrent())

		assert.True(t, iterator.MoveNext())
		assert.Equal(t, 2, iterator.GetCurrent())

		assert.False(t, iterator.MoveNext())
	})
}
