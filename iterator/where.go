package iterator

type whereIterator[T any] struct {
	predicate     func(T) bool
	innerIterator Iterator[T]
}

func (iterator *whereIterator[T]) GetCurrent() T {
	return iterator.innerIterator.GetCurrent()
}

func (iterator *whereIterator[T]) MoveNext() bool {
	for iterator.innerIterator.MoveNext() {
		item := iterator.innerIterator.GetCurrent()
		if iterator.predicate(item) {
			return true
		}
	}
	return false
}

func (iterator *whereIterator[T]) Clone() Iterator[T] {
	return NewWhereIterator(iterator.innerIterator.Clone(), iterator.predicate)
}

func NewWhereIterator[T any](innerIterator Iterator[T], predicate func(T) bool) Iterator[T] {
	iterator := whereIterator[T]{
		predicate:     predicate,
		innerIterator: innerIterator,
	}

	return &iterator
}
