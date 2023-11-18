package iterator

type Iterator[T any] interface {
	// Next advances to the next element if such an element exists.
	// It returns true iff the next element exists.
	Next() bool

	// Value returns the current value of the iterator.
	// Invoking Value on a drained iterator should return the zero value of T.
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
