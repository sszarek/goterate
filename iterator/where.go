package iterator

type whereIterator[T any] struct {
	predicate     func(T) bool
	innerIterator Iterator[T]
	current       T
}

func (iterator *whereIterator[T]) GetCurrent() T {
	return iterator.current
}

func (iterator *whereIterator[T]) HasNext() bool {
	for iterator.innerIterator.MoveNext() {
		item := iterator.innerIterator.GetCurrent()
		if iterator.predicate(item) {
			return true
		}
	}
	return false
}

func newWhereIterator[T any](innerIterator Iterator[T], predicate func(T) bool) *whereIterator[T] {
	iterator := whereIterator[T]{
		predicate: predicate,
	}

	return &iterator
}
