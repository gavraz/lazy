package iterator

type Iterator[T any] interface {
	Next() bool
	Value() T
}

type SizedIterator[T any] struct {
	Iterator[T]
	size int
}

func (s SizedIterator[T]) Size() int {
	return s.size
}

type FuncIterator[T any] struct {
	v T
	f func() (T, bool)
}

func (f *FuncIterator[T]) Next() bool {
	v, ok := f.f()
	if !ok {
		return false
	}

	f.v = v
	return true
}

func (f *FuncIterator[T]) Value() T {
	return f.v
}
