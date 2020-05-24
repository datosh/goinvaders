package vec2

import "image"

type Rect struct {
	Min, Max T
}

func NewRect(x, y, w, h float64) *Rect {
	return &Rect{
		Min: T{X: x, Y: y},
		Max: T{X: x + w, Y: y + h},
	}
}

func (r *Rect) X() float64 {
	return r.Min.X
}

func (r *Rect) Y() float64 {
	return r.Min.Y
}

func (r *Rect) Width() float64 {
	return r.Max.X - r.Min.X
}

func (r *Rect) Height() float64 {
	return r.Max.Y - r.Min.Y
}

func (r *Rect) Intersects(other *Rect) bool {
	// AABB collision test
	return r.Min.X < other.Max.X && other.Min.X < r.Max.X &&
		r.Min.Y < other.Max.Y && other.Min.Y < r.Max.Y
}

func (r *Rect) Inside(other *Rect) bool {
	return other.Min.X <= r.Min.X && r.Max.X <= other.Max.X &&
		other.Min.Y <= r.Min.Y && r.Max.Y <= other.Max.Y
}

func (r *Rect) FromImageRect(from image.Rectangle) {
	min := from.Min
	r.Min = T{float64(min.X), float64(min.Y)}
	max := from.Max
	r.Max = T{float64(max.X), float64(max.Y)}
}
