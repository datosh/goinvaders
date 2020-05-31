// Package vec2 provides two dimensional vector code.
package vec2

import "math"

// T represents a two dimensonal vector based on float64.
// Imported by other packages as vec2.T, therefore we don't repeat type name.
type T struct {
	X, Y float64
}

// I represents a two dimensonal vector based on integers.
type I struct {
	X, Y int
}

// New creates a new *T from two coordinates
func New(x, y float64) *T {
	return &T{X: x, Y: y}
}

// NewI creates a new *I from two coordinates
func NewI(x, y int) *I {
	return &I{X: x, Y: y}
}

// UX is unit vector where X=1
func UX() *T {
	return &T{1.0, 0.0}
}

// UY is unit vector where Y=1
func UY() *T {
	return &T{0.0, 1.0}
}

// UXY is unit vector where X=1 and Y=1
func UXY() *T {
	return &T{1.0, 1.0}
}

// Coords returns the single cooridnates of vector
func (t *T) Coords() (float64, float64) {
	return t.X, t.Y
}

func (t *T) Copy() *T {
	p := *t
	return &p
}

func (t *T) Null() bool {
	return t.X == 0 && t.Y == 0
}

// Add other to receiver. Other is unchanged.
// Receiver is returned for easy chaning.
func (t *T) Add(other *T) *T {
	t.X += other.X
	t.Y += other.Y
	return t
}

// Added returns a new *T which is the sum of receiver and other.
func (t *T) Added(other *T) *T {
	p := *t
	return p.Add(other)
}

// Sub subtracts other from receiver. Other is unchanged.
// Receiver is returned for easy chaning.
func (t *T) Sub(other *T) *T {
	t.X -= other.X
	t.Y -= other.Y
	return t
}

// Subed returns a new *T which is the difference of receiver and other.
func (t *T) Subed(other *T) *T {
	p := *t
	return p.Sub(other)
}

// Mul multiplies the scalar to receiver.
// Receiver is returned for easy chaning.
func (t *T) Mul(scalar float64) *T {
	t.X *= scalar
	t.Y *= scalar
	return t
}

// Muled returns a new *T which is receiver multiplied by scalar.
func (t *T) Muled(scalar float64) *T {
	p := *t
	return p.Mul(scalar)
}

// Invert vector.
// Receiver is returned for easy chaning.
func (t *T) Invert() *T {
	t.X = -t.X
	t.Y = -t.Y
	return t
}

// Inverted returns a new *T which is receiver inverted.
func (t *T) Inverted() *T {
	p := *t
	return p.Invert()
}

// Length returns length of receiver.
func (t *T) Length() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y)
}

// Normalize receiver. After calling Normalize, Length is always 1.
func (t *T) Normalize() {
	if t.Length() != 0 {
		t.Mul(1 / t.Length())
	}
}

// AsI converts a float64 based vector to an integer based one.
func (t *T) AsI() *I {
	return &I{int(t.X), int(t.Y)}
}

// AsT converts an integer based vector to a float64 based one.
func (i *I) AsT() *T {
	return &T{X: float64(i.X), Y: float64(i.Y)}
}
