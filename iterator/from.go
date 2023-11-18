package iterator

import "golang.org/x/exp/constraints"

func none[T any]() (T, bool) {
	var v T
	return v, false
}

func FromFunc[T any](next func() (T, bool)) Iterator[T] {
	return &FuncIterator[T]{f: next}
}

func ByDelta[T constraints.Ordered](begin, end, delta T) Iterator[T] {
	curr := begin
	return FromFunc(func() (T, bool) {
		next := curr + delta
		if next > end {
			return none[T]()
		}
		v := curr
		curr = next
		return v, true
	})
}

// TODO: we can create:
// 		From(x).To(y).By(d)
// 		From(x).To(y)
//		To(y).By(d)

func To(end int) Iterator[int] {
	return ByDelta[int](0, end, 1)
}

func Generate[T any](rnd func() T) Iterator[T] {
	return FromFunc(func() (T, bool) {
		return rnd(), true
	})
}

func FromSlice[T any](data []T) Iterator[T] {
	i := 0
	return FromFunc(func() (T, bool) {
		if i >= len(data) {
			return none[T]()
		}

		v := data[i]
		i++
		return v, true
	})
}

func FromValues[T any](v ...T) Iterator[T] {
	return FromSlice(v)
}
