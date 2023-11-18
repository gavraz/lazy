package iterator

import "golang.org/x/exp/constraints"

func none[T any]() (T, bool) {
	var v T
	return v, false
}

func FromIterFunc[T any](next func() (T, bool)) Iterator[T] {
	return &FuncIterator[T]{f: next}
}

func FromDelta[T constraints.Ordered](begin, end, delta T) Iterator[T] {
	curr := begin
	return FromIterFunc(func() (T, bool) {
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
	return FromDelta[int](0, end, 1)
}

func FromGenerator[T any](rnd func() T) Iterator[T] {
	return FromIterFunc(func() (T, bool) {
		return rnd(), true
	})
}

func FromSlice[T any](data []T) Iterator[T] {
	i := 0
	return FromIterFunc(func() (T, bool) {
		if i >= len(data) {
			return none[T]()
		}

		i++
		return data[i], true
	})
}

func FromValues[T any](v ...T) Iterator[T] {
	return FromSlice(v)
}
