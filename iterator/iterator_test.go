package iterator

import (
	"fmt"
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

		assert.True(t, iter.HasNext())
		iter.GetNext()
		assert.False(t, iter.HasNext())
	})
}

func TestGetNext(t *testing.T) {
	t.Run("Empty integer slice", func(t *testing.T) {
		slice := []int{}
		iter := NewIterator(slice)

		err, _ := iter.GetNext()

		assert.Error(t, err, "No more elements to iterate throught")
	})

	t.Run("Integer slice with 3 elements", func(t *testing.T) {
		slice := []int{1, 2, 3}
		iter := NewIterator(slice)

		_, first := iter.GetNext()
		assert.Equal(t, 1, first)
		_, second := iter.GetNext()
		assert.Equal(t, 2, second)
		_, third := iter.GetNext()
		assert.Equal(t, 3, third)
	})
}

func TestResult(t *testing.T) {
	t.Run("Empty integer slice", func(t *testing.T) {
		slice := []int{}
		iter := NewIterator(slice)

		result := iter.Result()
		assert.ElementsMatch(t, result, slice)
	})

	t.Run("Non-empty integer slice", func(t *testing.T) {
		slice := []int{1, 2, 3, 4}
		iter := NewIterator(slice)

		result := iter.Result()
		assert.ElementsMatch(t, result, slice)
	})
}

func TestTake(t *testing.T) {
	t.Run("Empty integer slice", func(t *testing.T) {
		slice := []int{}
		iter := NewIterator(slice)
		result := iter.Take(2).Result()

		assert.ElementsMatch(t, result, []int{})
	})

	t.Run("4 element integer slice", func(t *testing.T) {
		slice := []int{1, 2, 3, 4}
		iter := NewIterator(slice)
		result := iter.Take(2).Result()

		assert.ElementsMatch(t, result, []int{1, 2})
	})

	t.Run("4 element integer slice, negative param", func(t *testing.T) {
		slice := []int{1, 2, 3, 4}
		iter := NewIterator(slice)

		assert.PanicsWithError(t, "Non-negative integer expected as a parameter, got -2", func() {
			iter.Take(-2).Result()
		})
	})
}

func TestSkip(t *testing.T) {
	var testCases = []struct {
		skip     int
		input    []int
		expected []int
	}{
		{1, []int{}, []int{}},
		{1, []int{1, 2, 3}, []int{2, 3}},
		{2, []int{1, 2, 3, 4}, []int{3, 4}},
	}

	for _, tt := range testCases {
		name := fmt.Sprintf("%v,%v,%v", tt.skip, tt.input, tt.expected)

		t.Run(name, func(t *testing.T) {
			iter := NewIterator(tt.input)
			actual := iter.Skip(tt.skip).Result()

			assert.ElementsMatch(t, actual, tt.expected)
		})
	}

	t.Run("negative param provided", func(t *testing.T) {
		slice := []int{1, 2, 3, 4}
		iter := NewIterator(slice)

		assert.PanicsWithError(t, "Non-negative integer expected as a parameter, got -2", func() {
			iter.Skip(-2).Result()
		})
	})
}
