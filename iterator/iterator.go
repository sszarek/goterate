package iterator

type Iterator[T any] interface {
	GetCurrent() T
	MoveNext() bool
	Clone() Iterator[T]
}

type sliceIterator[T any] struct {
	curIdx int
	slice  []T
}

func (iter *sliceIterator[T]) GetCurrent() T {
	return iter.slice[iter.curIdx]
}

func (iter *sliceIterator[T]) MoveNext() bool {
	if len(iter.slice) <= iter.curIdx+1 {
		return false
	}

	iter.curIdx += 1
	return true
}

func (iter *sliceIterator[T]) Clone() Iterator[T] {
	return NewSliceIterator(iter.slice)
}

func NewSliceIterator[T any](slice []T) Iterator[T] {
	iter := sliceIterator[T]{
		slice:  slice,
		curIdx: 0,
	}
	return &iter
}
