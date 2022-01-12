package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMethodCombinations(t *testing.T) {
	t.Run("slice -> take -> where", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		iter := NewIterator(input)
		result := iter.
			Take(3).
			Where(func(t int) bool { return t > 2 }).
			Result()

		assert.ElementsMatch(t, result, []int{3})
	})

	t.Run("slice -> take -> where -> take", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		iter := NewIterator(input)
		result := iter.
			Take(5).
			Where(func(t int) bool { return t > 2 }).
			Take(2).
			Result()

		assert.ElementsMatch(t, result, []int{3, 4})
	})

	t.Run("slice -> where -> where -> where", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		iter := NewIterator(input)
		result := iter.
			Where(func(t int) bool { return t > 2 }).
			Where(func(t int) bool { return t < 8 }).
			Where(func(t int) bool { return t > 5 }).
			Result()

		assert.ElementsMatch(t, result, []int{6, 7})
	})
}
