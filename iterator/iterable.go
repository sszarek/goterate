package iterator

type Iterable[T any] interface {
	Where(predicate func(T) bool) Iterable[T]
	GetIterator() Iterator[T]
}

type sliceIterable[T any] struct {
	slice    []T
	iterator Iterator[T]
}

func (iterable sliceIterable[T]) GetIterator() Iterator[T] {
	return iterable.iterator.Clone()
}

func (iterable sliceIterable[T]) Where(predicate func(T) bool) Iterable[T] {
	whereIterator := NewWhereIterator(iterable.iterator, predicate)

	iterable.iterator = whereIterator

	// TODO: Change this to return new iterable
	return iterable
}

func NewSliceIterable[T any](slice []T) Iterable[T] {
	iterable := sliceIterable[T]{
		slice:    slice,
		iterator: NewSliceIterator(slice),
	}

	return &iterable
}
