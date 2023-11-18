package iterator

func FromIterFunc[T any](next func() (T, bool)) Iterator[T] {
	return &FuncIterator[T]{f: next}
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
			var v T
			return v, false
		}

		i++
		return data[i], true
	})
}
