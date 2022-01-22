package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimplecases(t *testing.T) {
	var result int
	iter := &whereIterator[int]{
		predicate: func(val int) bool { return val > 0 },
		current:   result,
	}

	assert.NotNil(t, iter)
}
