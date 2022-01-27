package iterator

type takeIterator[T any] struct {
	innerIterator Iterator[T]
	take          int
}

func (iterator *takeIterator[T]) GetCurrent() T {
	return iterator.innerIterator.GetCurrent()
}
