package iterator

import (
	"fmt"
)

type Iterator[T any] interface {
	GetNext() (error, T)
	HasNext() bool
	Where(predicate func(T) bool) Iterator[T]
	Take(count int) Iterator[T]
	Skip(count int) Iterator[T]
	First() (error, T)
	Last() (error, T)
	Result() []T
}

type iterImpl[T any] struct {
	curIdx int
	slice  []T
}

func (iter *iterImpl[T]) GetNext() (error, T) {
	iter.curIdx++

	var result T
	if iter.curIdx >= len(iter.slice) {
		return fmt.Errorf("No more elements to iterate throught"), result
	}

	result = iter.slice[iter.curIdx]

	return nil, result
}

func (iter *iterImpl[T]) HasNext() bool {
	return iter.curIdx < (len(iter.slice) - 1)
}

func (iter *iterImpl[T]) Where(predicate func(T) bool) Iterator[T] {
	result := []T{}

	for _, item := range iter.slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return NewIterator(result)
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
	if count < 0 {
		panic(fmt.Errorf("Non-negative integer expected as a parameter, got %v", count))
	}

	minEl := count
	if count > len(iter.slice) {
		minEl = len(iter.slice)
	}
	return NewIterator(iter.slice[minEl:])
}

func (iter *iterImpl[T]) Result() []T {
	return iter.slice
}

func (iter *iterImpl[T]) First() (error, T) {
	var result T
	if len(iter.slice) == 0 {
		return fmt.Errorf("Iterator is empty"), result
	}

	result = iter.slice[0]

	return nil, result
}

func (iter *iterImpl[T]) Last() (error, T) {
	var result T
	if len(iter.slice) == 0 {
		return fmt.Errorf("Iterator is empty"), result
	}

	result = iter.slice[len(iter.slice)-1]

	return nil, result
}

func NewIterator[T any](slice []T) Iterator[T] {
	iter := iterImpl[T]{
		slice:  slice,
		curIdx: -1,
	}
	return &iter
}
