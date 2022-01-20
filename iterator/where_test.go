package iterator

import "testing"

func TestSimplecases(t *testing.T) {
	var default int
	iter := &whereIterator[int]{
		predicate: func(val int) bool {return t > 0},
		current: default,
		
	}

	iter.
}
