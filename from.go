package lazy

func none[T any]() (T, bool) {
	var v T
	return v, false
}

func FromFunc[T any](next func() (T, bool)) Iterator[T] {
	return Iterator[T]{f: next, size: unknown}
}

func Generate[T any](rnd func() T) Iterator[T] {
	return FromFunc(func() (T, bool) {
		return rnd(), true
	})
}

func FromSlice[T any](data []T) Iterator[T] {
	i := 0
	it := FromFunc(func() (T, bool) {
		if i >= len(data) {
			return none[T]()
		}

		v := data[i]
		i++
		return v, true
	})
	it.size = len(data)
	return it
}

func FromValues[T any](v ...T) Iterator[T] {
	return FromSlice(v)
}
