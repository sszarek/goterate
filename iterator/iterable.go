package iterator

type Iterable[T any] interface {
	Where(predicate func(T) bool) Iterable[T]
	Take(take int) Iterable[T]
	GetIterator() Iterator[T]
}

type sliceIterable[T any] struct {
	slice    []T
	iterator Iterator[T]
}

func (iterable *sliceIterable[T]) GetIterator() Iterator[T] {
	return iterable.iterator.Clone()
}

func (iterable *sliceIterable[T]) Where(predicate func(T) bool) Iterable[T] {
	whereIterator := NewWhereIterator(iterable.iterator, predicate)

	newIterable := &sliceIterable[T]{
		slice:    iterable.slice,
		iterator: whereIterator,
	}

	return newIterable
}

func (iterable *sliceIterable[T]) Take(take int) Iterable[T] {
	takeIterator := NewTakeIterator(iterable.iterator, take)

	newIterable := &sliceIterable[T]{
		slice:    iterable.slice,
		iterator: takeIterator,
	}

	return newIterable
}

func NewSliceIterable[T any](slice []T) Iterable[T] {
	iterable := sliceIterable[T]{
		slice:    slice,
		iterator: NewSliceIterator(slice),
	}

	return &iterable
}
