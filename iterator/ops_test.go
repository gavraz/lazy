package iterator

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func Test_Limit(t *testing.T) {
	ten := To(10)
	r := To(100)
	r = Limit(r, 10)

	assert.True(t, slices.Equal(Slice(ten), Slice(r)))
}

func Test_Filter(t *testing.T) {
	even := Filter(To(10), func(x int) bool {
		return x%2 == 0
	})
	d2 := ByDelta(0, 10, 2)

	assert.True(t, slices.Equal(Slice(d2), Slice(even)))
}

func Test_Discard(t *testing.T) {
	d := Discard(FromValues(1, 2, 3, 4, 5), 3)
	assert.True(t, slices.Equal([]int{4, 5}, Slice(d)))

	d = Discard(FromValues(1, 2, 3), 10)
	assert.True(t, slices.Equal([]int{}, Slice(d)))
}

func Test_Paginate(t *testing.T) {
	p := Paginate(FromValues(1, 2, 3), 1, 3)
	assert.True(t, slices.Equal([]int{1, 2, 3}, Slice(p)))

	p = Paginate(FromValues(1, 2, 3, 3, 4, 5, 6, 7), 2, 3)
	assert.True(t, slices.Equal([]int{3, 4, 5}, Slice(p)))
}
