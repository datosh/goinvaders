package vec2

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRect(t *testing.T) {
	t.Run("New from dimensions", func(t *testing.T) {
		r := NewRect(10.0, 20.0, 30.0, 40.0)

		assert.Equal(t, 10.0, r.X())
		assert.Equal(t, 20.0, r.Y())
		assert.Equal(t, 30.0, r.Width())
		assert.Equal(t, 40.0, r.Height())
	})

	t.Run("New from points", func(t *testing.T) {
		p1 := T{10, 10}
		p2 := T{20, 20}
		r := Rect{p1, p2}

		assert.Equal(t, 10.0, r.X())
		assert.Equal(t, 10.0, r.Y())
		assert.Equal(t, 10.0, r.Width())
		assert.Equal(t, 10.0, r.Height())
	})
}

func TestRect_Intersects(t *testing.T) {
	r1 := NewRect(10, 10, 10, 10)
	r2 := NewRect(10, 10, 10, 10)
	r3 := NewRect(20, 20, 10, 10)
	r4 := NewRect(15, 15, 10, 10)

	t.Run("same intersects", func(t *testing.T) {
		assert.True(t, r1.Intersects(r2))
	})

	t.Run("same intersects - order", func(t *testing.T) {
		assert.True(t, r2.Intersects(r1))
	})

	t.Run("same intersects - same object", func(t *testing.T) {
		assert.True(t, r1.Intersects(r1))
	})

	t.Run("touching doesnt intersect", func(t *testing.T) {
		assert.False(t, r1.Intersects(r3))
	})

	t.Run("intersects", func(t *testing.T) {
		assert.True(t, r1.Intersects(r4))
	})
}

func TestRect_Inside(t *testing.T) {
	r1 := NewRect(10, 10, 10, 10)
	r2 := NewRect(12, 12, 5, 5)
	r3 := NewRect(15, 15, 10, 2)

	t.Run("inside", func(t *testing.T) {
		assert.True(t, r2.Inside(r1))
	})

	t.Run("same is also inside", func(t *testing.T) {
		assert.True(t, r1.Inside(r1))
	})

	t.Run("full overlay", func(t *testing.T) {
		assert.False(t, r1.Inside(r2))
	})

	t.Run("exit right", func(t *testing.T) {
		assert.False(t, r3.Inside(r1))
	})
}

func TestRect_FromImageRect(t *testing.T) {
	imgRect := image.Rect(10, 10, 20, 20)
	r := Rect{}
	r.FromImageRect(imgRect)

	assert.Equal(t, 10.0, r.X())
	assert.Equal(t, 10.0, r.Y())
	assert.Equal(t, 10.0, r.Width())
	assert.Equal(t, 10.0, r.Height())
}
