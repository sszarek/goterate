package iterator

type Iterable[T any] interface {
	Where(predicate func(T) bool) Iterable[T]
	GetIterator() Iterator[T]
}

type sliceIterable[T any] struct {
	slice  []T
	curIdx int
}

func (iterable sliceIterable[T]) GetIterator() Iterator[T] {
	return NewSliceIterator(iterable.slice)
}

func (iterable sliceIterable[T]) Where(predicate func(T) bool) Iterable[T] {
	return nil
}

func NewSliceIterable[T any](slice []T) Iterable[T] {
	iterable := sliceIterable[T]{
		slice: slice,
	}

	return &iterable
}
