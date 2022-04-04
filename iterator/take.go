package iterator

import "fmt"

type takeIterator[T any] struct {
	innerIterator Iterator[T]
	take          int
	curIdx        int
}

func (iterator *takeIterator[T]) GetCurrent() T {
	return iterator.innerIterator.GetCurrent()
}

func (iterator *takeIterator[T]) MoveNext() bool {
	if iterator.curIdx+1 < iterator.take {
		iterator.curIdx++
		return iterator.innerIterator.MoveNext()
	}

	return false
}

func (iterator *takeIterator[T]) Clone() Iterator[T] {
	return NewTakeIterator(iterator.innerIterator, iterator.take)
}

func NewTakeIterator[T any](innerIterator Iterator[T], take int) Iterator[T] {
	if innerIterator == nil {
		panicMsg := "Expected 'innerIter' to be Iterator but received 'nil'"
		panic(panicMsg)
	}

	if take < 0 {
		panicMsg := fmt.Sprintf("Expected 'take' to be non-negative number but received: %d", take)
		panic(panicMsg)
	}

	iterator := takeIterator[T]{
		innerIterator: innerIterator,
		take:          take,
		curIdx:        -1,
	}

	return &iterator
}
