package lazy

import (
	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Integer | constraints.Float
}

// Range returns an iterator for a range [begin, begin+delta, ..., end).
// Values are assumed to be non-negative.
// Note: end=0 represents infinity.
func Range[T Numeric](begin, end, delta T) Iterator[T] {
	curr := begin
	return FromFunc(func() (T, bool) {
		next := curr + delta
		if inf := end == 0; next > end && !inf {
			return none[T]()
		}
		v := curr
		curr = next
		return v, true
	})
}

type Ranger[T Numeric] struct {
	Iterator[T]
	begin T
	end   T
	delta T
}

func (r Ranger[T]) To(end T) Ranger[T] {
	r.end = end
	r.Iterator = Range[T](r.begin, r.end, r.delta)
	return r
}

func (r Ranger[T]) By(delta T) Iterator[T] {
	cloned := r
	cloned.delta = delta
	cloned.Iterator = Range[T](cloned.begin, cloned.end, cloned.delta)
	return cloned.Iterator
}

func From[T Numeric](begin T) Ranger[T] {
	r := Ranger[T]{begin: begin, end: 0, delta: 1}
	r.Iterator = Range[T](r.begin, r.end, r.delta)
	return r
}

// To returns an iterator for the range: [0, end)
func To(end int) Iterator[int] {
	return Range[int](0, end, 1)
}
