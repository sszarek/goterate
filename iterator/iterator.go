package iterator

import "fmt"

type Iterator[T any] interface {
	GetNext() T
	HasNext() bool
}

type iterImpl[T any] struct {
	curIdx int
	slice  []T
}

func (iter *iterImpl[T]) GetNext() T {
	iter.curIdx++
	if iter.curIdx >= len(iter.slice) {
		panic(fmt.Errorf("No more elements to iterate throught"))
	}
	return iter.slice[iter.curIdx]
}

func (iter *iterImpl[T]) HasNext() bool {
	return len(iter.slice) > 0
}

func NewIterator[T any](slice []T) Iterator[T] {
	iter := iterImpl[T]{
		slice:  slice,
		curIdx: -1,
	}
	return &iter
}
