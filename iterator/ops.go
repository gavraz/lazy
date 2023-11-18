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

func Discard[T any](iter Iterator[T], count int) Iterator[T] {
	i := 0
	return FromFunc(func() (T, bool) {
		for ; i < count && iter.Next(); i++ {
		}

		ok := iter.Next()
		return iter.Value(), ok
	})
}

func Paginate[T any](iter Iterator[T], page, count int) Iterator[T] {
	skipped := Discard(iter, (page-1)*count)
	return Limit(skipped, count)
}

type Sizer interface {
	Size() int
}

func Slice[T any](iter Iterator[T]) []T {
	var all []T

	sized, ok := iter.(Sizer)
	if ok {
		all = make([]T, 0, sized.Size())
	}

	for iter.Next() {
		all = append(all, iter.Value())
	}

	return all
}
