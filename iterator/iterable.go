package iterator

type Iterable[T any] interface {
	Filter(predicate func(T) bool) Iterable[T]
	Take(take int) Iterable[T]
	GetIterator() Iterator[T]
}

type iterableImpl[T any] struct {
	iterator Iterator[T]
}

func (iterable *iterableImpl[T]) GetIterator() Iterator[T] {
	return iterable.iterator.Clone()
}

func (iterable *iterableImpl[T]) Filter(predicate func(T) bool) Iterable[T] {
	filterIterator := NewFilterIterator(iterable.iterator, predicate)
	return iterable.newFromIterator(filterIterator)
}

func (iterable *iterableImpl[T]) Take(take int) Iterable[T] {
	takeIterator := NewTakeIterator(iterable.iterator, take)
	return iterable.newFromIterator(takeIterator)
}

func NewSliceIterable[T any](slice []T) Iterable[T] {
	iterable := iterableImpl[T]{
		iterator: NewSliceIterator(slice),
	}

	return &iterable
}

func (iterable *iterableImpl[T]) newFromIterator(iterator Iterator[T]) Iterable[T] {
	newIterable := &iterableImpl[T]{
		iterator: iterator,
	}

	return newIterable
}
