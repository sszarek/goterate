package iterator

import (
	"fmt"
)

type Iterator[T any] interface {
	GetCurrent() T
	MoveNext() bool
}

type sliceIterator[T any] struct {
	curIdx int
	slice  []T
}

func (iter *sliceIterator[T]) GetNext() (error, T) {
	iter.curIdx++

	var result T
	if iter.curIdx >= len(iter.slice) {
		return fmt.Errorf("No more elements to iterate throught"), result
	}

	result = iter.slice[iter.curIdx]

	return nil, result
}

func (iter *sliceIterator[T]) HasNext() bool {
	return iter.curIdx < (len(iter.slice) - 1)
}

func NewSliceIterator[T any](slice []T) Iterator[T] {
	iter := sliceIterator[T]{
		slice:  slice,
		curIdx: -1,
	}
	return &iter
}
