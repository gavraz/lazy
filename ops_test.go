package lazy

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"slices"
	"strings"
	"testing"
)

func Test_Limit(t *testing.T) {
	ten := To(10)
	r := To(100).Limit(10)

	limited := r.Slice()
	assert.True(t, slices.Equal(ten.Slice(), limited))

	// verify the optimization works with exact allocation, 16 otherwise
	assert.Equal(t, 10, cap(limited))
}

func Test_Filter(t *testing.T) {
	even := To(10).Filter(func(x int) bool {
		return x%2 == 0
	})
	d2 := Range(0, 10, 2)

	assert.True(t, slices.Equal(d2.Slice(), even.Slice()))
}

func Test_Discard(t *testing.T) {
	d := FromValues(1, 2, 3, 4, 5).Discard(3)
	assert.True(t, slices.Equal([]int{4, 5}, d.Slice()))

	d = FromValues(1, 2, 3).Discard(10)
	assert.True(t, slices.Equal([]int{}, d.Slice()))
}

func Test_Paginate(t *testing.T) {
	p := FromValues(1, 2, 3).Paginate(1, 3)
	assert.True(t, slices.Equal([]int{1, 2, 3}, p.Slice()))

	p = FromValues(1, 2, 3, 3, 4, 5, 6, 7).Paginate(2, 3)
	assert.True(t, slices.Equal([]int{3, 4, 5}, p.Slice()))
}

func Test_Map(t *testing.T) {
	m := Map(FromValues(1, 2, 3), func(x int) string {
		return fmt.Sprint(x)
	})

	res := m.Slice()
	exp := []string{"1", "2", "3"}
	assert.Equal(t, len(exp), len(res))
	for i, s := range exp {
		strings.Compare(s, res[i])
	}
}
