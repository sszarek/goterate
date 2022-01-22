package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIteration(t *testing.T) {
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
