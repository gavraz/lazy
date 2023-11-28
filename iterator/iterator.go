package iterator

type Easy[T any] struct {
	v T
	f Iterator[T]
}

func (ez *Easy[T]) Next() bool {
	v, ok := ez.f.f()
	if !ok {
		return false
	}

	ez.v = v
	return true
}

func (ez *Easy[T]) Value() T {
	return ez.v
}
