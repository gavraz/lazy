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
