package iterator

func none[T any]() (T, bool) {
	var v T
	return v, false
}

func FromFunc[T any](next func() (T, bool)) Iterator[T] {
	return &FuncIterator[T]{f: next}
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
