package lazy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Generate(t *testing.T) {
	k := 2
	g := Generate[int](func() int {
		if k > 0 {
			k--
			return 1
		}

		return 2
	}).Easy()

	assert.True(t, g.Next())
	assert.Equal(t, 1, g.Value())
	assert.True(t, g.Next())
	assert.Equal(t, 1, g.Value())

	assert.True(t, g.Next())
	assert.Equal(t, 2, g.Value())
	assert.True(t, g.Next())
	assert.Equal(t, 2, g.Value())
}

func Test_FromSlice(t *testing.T) {
	slice := FromSlice([]int{}).Easy()
	assert.False(t, slice.Next())
}

func Test_FromValues(t *testing.T) {
	vals := FromValues(1).Easy()
	assert.True(t, vals.Next())
	assert.Equal(t, 1, vals.Value())
	assert.False(t, vals.Next())

	vals = FromValues(1, 3, 5).Easy()
	assert.True(t, vals.Next())
	assert.Equal(t, 1, vals.Value())
	assert.True(t, vals.Next())
	assert.Equal(t, 3, vals.Value())
	assert.True(t, vals.Next())
	assert.Equal(t, 5, vals.Value())
	assert.False(t, vals.Next())
}
