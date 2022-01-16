package iterator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAny(t *testing.T) {
	var testCases = []struct {
		predicate func(int) bool
		input     []int
		expected  bool
	}{
		{
			func(num int) bool {
				return true
			}, []int{}, false,
		},
		{
			func(num int) bool {
				return false
			}, []int{}, false,
		},
		{
			func(num int) bool {
				return num > 0
			}, []int{-1}, false},
		{
			func(num int) bool {
				return num > 0
			}, []int{1, 2, 3}, true,
		},
		{
			func(num int) bool {
				return num < 2
			}, []int{1, 2, 3}, true,
		},
	}

	for _, tt := range testCases {
		name := fmt.Sprintf("%v,%v", tt.input, tt.expected)

		t.Run(name, func(t *testing.T) {
			iter := NewIterator(tt.input)
			actual := iter.Any(tt.predicate)

			assert.Equal(t, actual, tt.expected)
		})
	}
}

func TestAll(t *testing.T) {
	var testCases = []struct {
		predicate func(int) bool
		input     []int
		expected  bool
	}{
		{
			func(num int) bool {
				return true
			}, []int{}, true,
		},
		{
			func(num int) bool {
				return false
			}, []int{}, true,
		},
		{
			func(num int) bool {
				return num > 0
			}, []int{-1}, false},
		{
			func(num int) bool {
				return num > 0
			}, []int{1, 2, 3}, true,
		},
		{
			func(num int) bool {
				return num < 2
			}, []int{1, 2, 3}, false,
		},
	}

	for _, tt := range testCases {
		name := fmt.Sprintf("%v,%v", tt.input, tt.expected)

		t.Run(name, func(t *testing.T) {
			iter := NewIterator(tt.input)
			actual := iter.All(tt.predicate)

			assert.Equal(t, actual, tt.expected)
		})
	}
}

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
	var testCases = []struct {
		take     int
		input    []int
		expected []int
	}{
		{1, []int{}, []int{}},
		{1, []int{1, 2, 3}, []int{1}},
		{2, []int{1, 2, 3, 4}, []int{1, 2}},
		{4, []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
	}

	for _, tt := range testCases {
		name := fmt.Sprintf("%v,%v,%v", tt.take, tt.input, tt.expected)

		t.Run(name, func(t *testing.T) {
			iter := NewIterator(tt.input)
			actual := iter.Take(tt.take).Result()

			assert.ElementsMatch(t, actual, tt.expected)
		})
	}

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
		{3, []int{1, 2, 3}, []int{}},
		{0, []int{1, 2, 3}, []int{1, 2, 3}},
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

func TestWhere(t *testing.T) {
	var testCases = []struct {
		predicate func(int) bool
		input     []int
		expected  []int
	}{
		{
			func(num int) bool {
				return num > 0
			}, []int{-1}, []int{}},
		{
			func(num int) bool {
				return num > 0
			}, []int{1, 2, 3}, []int{1, 2, 3},
		},
		{
			func(num int) bool {
				return num < 0
			}, []int{1, 2, 3}, []int{},
		},
	}

	for _, tt := range testCases {
		name := fmt.Sprintf("%v,%v", tt.input, tt.expected)

		t.Run(name, func(t *testing.T) {
			iter := NewIterator(tt.input)
			actual := iter.Where(tt.predicate).Result()

			assert.ElementsMatch(t, actual, tt.expected)
		})
	}
}

func TestFirst(t *testing.T) {
	t.Run("Empty int slice", func(t *testing.T) {
		input := []int{}
		iter := NewIterator(input)
		err, _ := iter.First()

		assert.Error(t, err, "Iterator is empty")
	})

	t.Run("One element int slice", func(t *testing.T) {
		input := []int{1}
		iter := NewIterator(input)
		err, actual := iter.First()

		assert.Nil(t, err)
		assert.Equal(t, 1, actual)
	})

	t.Run("Several elements int slice", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		iter := NewIterator(input)
		err, actual := iter.First()

		assert.Nil(t, err)
		assert.Equal(t, 1, actual)
	})
}

func TestLast(t *testing.T) {
	t.Run("Empty int slice", func(t *testing.T) {
		input := []int{}
		iter := NewIterator(input)
		err, _ := iter.Last()

		assert.Error(t, err, "Iterator is empty")
	})

	t.Run("One element int slice", func(t *testing.T) {
		input := []int{1}
		iter := NewIterator(input)
		err, actual := iter.Last()

		assert.Nil(t, err)
		assert.Equal(t, 1, actual)
	})

	t.Run("Several elements int slice", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		iter := NewIterator(input)
		err, actual := iter.Last()

		assert.Nil(t, err)
		assert.Equal(t, 5, actual)
	})
}
