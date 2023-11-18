package lazy

import "github.com/lazy/iterator"

type Op[T any] func(iterator iterator.Iterator[T]) iterator.Iterator[T]

func Limit[T any](limit int) Op[T] {
	return func(iter iterator.Iterator[T]) iterator.Iterator[T] {
		return iterator.Limit(iter, limit)
	}
}

func Map[T any](mp func(T) T) Op[T] {
	return func(iter iterator.Iterator[T]) iterator.Iterator[T] {
		return iterator.Map(iter, mp)
	}
}

func Filter[T any](filter func(T) bool) Op[T] {
	return func(iter iterator.Iterator[T]) iterator.Iterator[T] {
		return iterator.Filter(iter, filter)
	}
}

func Discard[T any](count int) Op[T] {
	return func(iter iterator.Iterator[T]) iterator.Iterator[T] {
		return iterator.Discard(iter, count)
	}
}

func Paginate[T any](page, count int) Op[T] {
	return func(iter iterator.Iterator[T]) iterator.Iterator[T] {
		return iterator.Paginate(iter, page, count)
	}
}

type SlicerIterator[T any] struct {
	iterator.Iterator[T]
}

func (s SlicerIterator[T]) Slice() []T {
	return iterator.Slice(s.Iterator)
}

func Build[T any](iter iterator.Iterator[T], ops ...Op[T]) SlicerIterator[T] {
	curr := iter
	for _, op := range ops {
		curr = op(curr)
	}

	return SlicerIterator[T]{Iterator: curr}
}
