package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIterator(t *testing.T) {
	t.Run("Nil slice", func(t *testing.T) {
		var slice []int
		slice = nil

		actual := NewIterator(slice)

		assert.NotNil(t, actual)
	})
}

func TestHasNext(t *testing.T) {
	t.Run("Empty slice", func(t *testing.T) {
		slice := []int{}
		iter := NewIterator(slice)

		actual := iter.HasNext()
		assert.False(t, actual)
	})

	t.Run("Non empty slice", func(t *testing.T) {
		slice := []int{1}
		iter := NewIterator(slice)

		actual := iter.HasNext()
		assert.True(t, actual)
	})
}

func TestGetNext(t *testing.T) {
	t.Run("Empty integer slice", func(t *testing.T) {
		slice := []int{}
		iter := NewIterator(slice)

		assert.PanicsWithError(t, "No more elements to iterate throught", func() {
			iter.GetNext()
		})
	})

	t.Run("Integer slice with 3 elements", func(t *testing.T) {
		slice := []int{1, 2, 3}
		iter := NewIterator(slice)
		assert.Equal(t, 1, iter.GetNext())
		assert.Equal(t, 2, iter.GetNext())
		assert.Equal(t, 3, iter.GetNext())
	})
}
