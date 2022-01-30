package iterator

type Iterable[T any] interface {
	Filter(predicate func(T) bool) Iterable[T]
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

func (iterable *sliceIterable[T]) Filter(predicate func(T) bool) Iterable[T] {
	filterIterator := NewFilterIterator(iterable.iterator, predicate)
	return iterable.newFromIterator(filterIterator)
}

func (iterable *sliceIterable[T]) Take(take int) Iterable[T] {
	takeIterator := NewTakeIterator(iterable.iterator, take)
	return iterable.newFromIterator(takeIterator)
}

func NewSliceIterable[T any](slice []T) Iterable[T] {
	iterable := sliceIterable[T]{
		slice:    slice,
		iterator: NewSliceIterator(slice),
	}

	return &iterable
}

func (iterable *sliceIterable[T]) newFromIterator(iterator Iterator[T]) Iterable[T] {
	newIterable := &sliceIterable[T]{
		slice:    iterable.slice,
		iterator: iterator,
	}

	return newIterable
}
