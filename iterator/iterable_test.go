package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSliceIterable(t *testing.T) {
	t.Run("creates iterator", func(t *testing.T) {
		slice := []int{1, 2, 3, 4}

		iterable := NewSliceIterable(slice)

		assert.NotNil(t, iterable)
	})
}

func TestGetIterator(t *testing.T) {
	t.Run("returns iterator", func(t *testing.T) {
		slice := []int{1, 2, 3, 4}

		iterable := NewSliceIterable(slice)
		iterator := iterable.GetIterator()

		assert.NotNil(t, iterator)
	})

	t.Run("clones iterator", func(t *testing.T) {
		slice := []int{1, 2, 3, 4}

		iterable := NewSliceIterable(slice)
		first := iterable.GetIterator()
		second := iterable.GetIterator()

		assert.NotSame(t, first, second)
	})
}
