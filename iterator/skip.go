package iterator

const startSkipValue = 0

type skipIterator[T any] struct {
	innerIterator Iterator[T]
	skip          int
	skipped       int
}

func (iter *skipIterator[T]) GetCurrent() T {
	return iter.innerIterator.GetCurrent()
}

func (iter *skipIterator[T]) MoveNext() bool {
	if iter.skipped < iter.skip {
		for iter.skipped != iter.skip {
			if iter.innerIterator.MoveNext() {
				iter.skipped++
			} else {
				return false
			}
		}
	}

	moveResult := iter.innerIterator.MoveNext()
	if moveResult {
		iter.skipped++
	}

	return moveResult
}

func (iter *skipIterator[T]) Clone() Iterator[T] {
	return NewTakeIterator(iter.innerIterator, iter.skip)
}

func NewSkipIterator[T any](innerIterator Iterator[T], skip int) Iterator[T] {
	iterator := &skipIterator[T]{
		innerIterator: innerIterator,
		skip:          skip,
		skipped:       startSkipValue,
	}
	return iterator
}
