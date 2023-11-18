package iterator

func Limit[T any](iter Iterator[T], limit int) Iterator[T] {
	i := 0
	limiter := FromFunc[T](func() (T, bool) {
		if i >= limit || !iter.Next() {
			return none[T]()
		}

		i++
		return iter.Value(), true
	})

	return &SizedIterator[T]{
		Iterator: limiter,
		size:     limit,
	}
}

func Filter[T any](iter Iterator[T], filter func(T) bool) Iterator[T] {
	return FromFunc(func() (T, bool) {
		for iter.Next() {
			v := iter.Value()
			if filter(v) {
				return v, true
			}
		}

		return none[T]()
	})
}

func Map[T any, S any](iter Iterator[T], m func(T) S) Iterator[S] {
	return FromFunc(func() (S, bool) {
		if !iter.Next() {
			return none[S]()
		}

		return m(iter.Value()), true
	})
}

func Slice[T any](iter Iterator[T]) []T {
	var all []T

	sized, ok := iter.(SizedIterator[T])
	if ok {
		all = make([]T, 0, sized.Size())
	}

	for iter.Next() {
		all = append(all, iter.Value())
	}

	return all
}

// TODO: paginate
