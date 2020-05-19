package vec2

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
	t.Run("Order of x & y", func(t *testing.T) {
		p := Point{2, 3}
		assert.Equal(t, p.X, 2.0)
		assert.Equal(t, p.Y, 3.0)
	})

	t.Run("Addition", func(t *testing.T) {
		a := Point{2, 3}
		b := Point{10, 20}

		a.Add(b)

		assert.Equal(t, a.X, 12.0)
		assert.Equal(t, a.Y, 23.0)
	})

	t.Run("Subtraction", func(t *testing.T) {
		a := Point{2, 20}
		b := Point{10, 3}

		a.Sub(b)

		assert.Equal(t, a.X, -8.0)
		assert.Equal(t, a.Y, 17.0)
	})

	t.Run("Multiplication", func(t *testing.T) {
		a := Point{2, 20}
		s := 3.0

		a.Mul(s)

		assert.Equal(t, a.X, 6.0)
		assert.Equal(t, a.Y, 60.0)
	})

	t.Run("Length", func(t *testing.T) {
		a := Point{1, 1}

		assert.Equal(t, a.Length(), math.Sqrt(2))
	})

	t.Run("Normalize", func(t *testing.T) {
		a := Point{10, 0}

		a.Normalize()

		assert.Equal(t, a.X, 1.0)
		assert.Equal(t, a.Y, 0.0)
	})

	t.Run("AsPointI", func(t *testing.T) {
		a := Point{1.5, 2.0}

		b := a.AsPointI()

		assert.Equal(t, b.X, 1)
		assert.Equal(t, b.Y, 2)
	})

	t.Run("AsPoint", func(t *testing.T) {
		a := PointI{1, 2}

		b := a.AsPoint()

		assert.Equal(t, b.X, 1.0)
		assert.Equal(t, b.Y, 2.0)
	})
}
