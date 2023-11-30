package lazy

const unknown = -1

type Iterator[T any] struct {
	f    func() (T, bool)
	size int
}

func (it Iterator[T]) Limit(limit int) Iterator[T] {
	i := 0
	limiter := FromFunc[T](func() (T, bool) {
		v, ok := it.f()
		if i >= limit || !ok {
			return none[T]()
		}

		i++
		return v, true
	})

	limiter.size = limit
	return limiter
}

func (it Iterator[T]) Filter(filter func(T) bool) Iterator[T] {
	return FromFunc(func() (T, bool) {
		for v, ok := it.f(); ok; v, ok = it.f() {
			if filter(v) {
				return v, true
			}
		}

		return none[T]()
	})
}

func (it Iterator[T]) Map(m func(T) T) Iterator[T] {
	return Map(it, m)
}

func Map[T any, S any](it Iterator[T], m func(T) S) Iterator[S] {
	mp := FromFunc(func() (S, bool) {
		v, ok := it.f()
		if !ok {
			return none[S]()
		}

		return m(v), true
	})
	mp.size = it.size
	return mp
}

func (it Iterator[T]) Discard(count int) Iterator[T] {
	i := 0
	return FromFunc(func() (T, bool) {
		v, ok := it.f()
		for ; i < count && ok; i++ {
			v, ok = it.f()
		}

		return v, ok
	})
}

func (it Iterator[T]) Paginate(page, count int) Iterator[T] {
	skipped := it.Discard((page - 1) * count)
	return skipped.Limit(count)
}

func (it Iterator[T]) Next() (T, bool) {
	return it.f()
}

func (it Iterator[T]) Easy() *Easy[T] {
	return &Easy[T]{f: it}
}

func (it Iterator[T]) Slice() []T {
	var all []T

	if it.size != unknown {
		all = make([]T, 0, it.size)
	}

	v, ok := it.f()
	for ok {
		all = append(all, v)
		v, ok = it.f()
	}

	return all
}
