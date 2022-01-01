package iterator

import (
	"fmt"
)

type Iterator[T any] interface {
	GetNext() T
	HasNext() bool
	Where(predicate func(T) bool) Iterator[T]
	Take(count int) Iterator[T]
	Skip(count int) Iterator[T]
	Result() []T
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
	return iter.curIdx < (len(iter.slice) - 1)
}

func (iter *iterImpl[T]) Where(predicate func(T) bool) Iterator[T] {
	return nil
}

func (iter *iterImpl[T]) Take(count int) Iterator[T] {
	if count < 0 {
		panic(fmt.Errorf("Non-negative integer expected as a parameter, got %v", count))
	}

	maxEl := count

	if count >= len(iter.slice) {
		maxEl = len(iter.slice)
	}
	return NewIterator(iter.slice[0:maxEl])
}

func (iter *iterImpl[T]) Skip(count int) Iterator[T] {
	return nil
}

func (iter *iterImpl[T]) Result() []T {
	return iter.slice
}

func NewIterator[T any](slice []T) Iterator[T] {
	iter := iterImpl[T]{
		slice:  slice,
		curIdx: -1,
	}
	return &iter
}
