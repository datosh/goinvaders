package vec2

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVec2(t *testing.T) {
	t.Run("Order of x & y", func(t *testing.T) {
		p := &T{2, 3}
		assert.Equal(t, p.X, 2.0)
		assert.Equal(t, p.Y, 3.0)
	})

	t.Run("Addition", func(t *testing.T) {
		a := &T{2, 3}
		b := &T{10, 20}

		a.Add(b)

		assert.Equal(t, a.X, 12.0)
		assert.Equal(t, a.Y, 23.0)
	})

	t.Run("Addition returns this", func(t *testing.T) {
		a := &T{2, 3}
		b := &T{10, 20}

		c := a.Add(b)

		assert.Same(t, c, a)
	})

	t.Run("Copy Addition", func(t *testing.T) {
		a := &T{2, 3}
		b := &T{10, 20}

		c := a.Added(b)

		assert.Equal(t, a.X, 2.0)
		assert.Equal(t, a.Y, 3.0)
		assert.Equal(t, b.X, 10.0)
		assert.Equal(t, b.Y, 20.0)
		assert.Equal(t, c.X, 12.0)
		assert.Equal(t, c.Y, 23.0)
	})

	t.Run("Subtraction", func(t *testing.T) {
		a := &T{2, 20}
		b := &T{10, 3}

		a.Sub(b)

		assert.Equal(t, a.X, -8.0)
		assert.Equal(t, a.Y, 17.0)
	})

	t.Run("Subtraction returns this", func(t *testing.T) {
		a := &T{2, 20}
		b := &T{10, 3}

		c := a.Sub(b)

		assert.Same(t, c, a)
	})

	t.Run("Copy Subtraction", func(t *testing.T) {
		a := &T{2, 20}
		b := &T{10, 3}

		c := a.Subed(b)

		assert.Equal(t, a.X, 2.0)
		assert.Equal(t, a.Y, 20.0)
		assert.Equal(t, b.X, 10.0)
		assert.Equal(t, b.Y, 3.0)
		assert.Equal(t, c.X, -8.0)
		assert.Equal(t, c.Y, 17.0)
	})

	t.Run("Multiplication", func(t *testing.T) {
		a := &T{2, 20}
		s := 3.0

		a.Mul(s)

		assert.Equal(t, a.X, 6.0)
		assert.Equal(t, a.Y, 60.0)
	})

	t.Run("Multiplication returns this", func(t *testing.T) {
		a := &T{2, 20}
		s := 3.0

		c := a.Mul(s)

		assert.Same(t, c, a)
	})

	t.Run("Copy Multiplication", func(t *testing.T) {
		a := &T{2, 20}
		s := 3.0

		c := a.Muled(s)

		assert.Equal(t, a.X, 2.0)
		assert.Equal(t, a.Y, 20.0)
		assert.Equal(t, s, 3.0)
		assert.Equal(t, c.X, 6.0)
		assert.Equal(t, c.Y, 60.0)
	})

	t.Run("Length of (1,1)=sqrt(2)", func(t *testing.T) {
		a := &T{1, 1}

		assert.Equal(t, math.Sqrt(2), a.Length())
	})

	t.Run("Length of (0,0)=0", func(t *testing.T) {
		a := &T{0, 0}

		assert.Equal(t, 0.0, a.Length())
	})

	t.Run("Normalize", func(t *testing.T) {
		a := &T{10, 0}

		a.Normalize()

		assert.Equal(t, a.X, 1.0)
		assert.Equal(t, a.Y, 0.0)
	})

	t.Run("AsI", func(t *testing.T) {
		a := &T{1.5, 2.0}

		b := a.AsI()

		assert.Equal(t, b.X, 1)
		assert.Equal(t, b.Y, 2)
	})

	t.Run("AsT", func(t *testing.T) {
		a := &I{1, 2}

		b := a.AsT()

		assert.Equal(t, b.X, 1.0)
		assert.Equal(t, b.Y, 2.0)
	})
}
