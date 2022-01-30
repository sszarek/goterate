package iterator

type filterIterator[T any] struct {
	predicate     func(T) bool
	innerIterator Iterator[T]
}

func (iterator *filterIterator[T]) GetCurrent() T {
	return iterator.innerIterator.GetCurrent()
}

func (iterator *filterIterator[T]) MoveNext() bool {
	for iterator.innerIterator.MoveNext() {
		item := iterator.innerIterator.GetCurrent()
		if iterator.predicate(item) {
			return true
		}
	}
	return false
}

func (iterator *filterIterator[T]) Clone() Iterator[T] {
	return NewFilterIterator(iterator.innerIterator.Clone(), iterator.predicate)
}

func NewFilterIterator[T any](innerIterator Iterator[T], predicate func(T) bool) Iterator[T] {
	iterator := filterIterator[T]{
		predicate:     predicate,
		innerIterator: innerIterator,
	}

	return &iterator
}
