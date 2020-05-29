package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Run("Filter even integers", func(t *testing.T) {
		a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

		b := Filter(a, func(i interface{}) bool { return i.(int)%2 == 0 }).([]int)

		assert.Equal(t, []int{0, 2, 4, 6, 8}, b)
	})
}
