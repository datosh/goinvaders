package vec2

import "image"

type Rect struct {
	Min, Max Point
}

func NewRect(x, y, w, h float64) *Rect {
	return &Rect{
		Min: Point{X: x, Y: y},
		Max: Point{X: x + w, Y: y + h},
	}
}

func (this *Rect) X() float64 {
	return this.Min.X
}

func (this *Rect) Y() float64 {
	return this.Min.Y
}

func (this *Rect) Width() float64 {
	return this.Max.X - this.Min.X
}

func (this *Rect) Height() float64 {
	return this.Max.Y - this.Min.Y
}

func (this *Rect) Intersects(other *Rect) bool {
	// AABB collision test
	return this.Min.X < other.Max.X && other.Min.X < this.Max.X &&
		this.Min.Y < other.Max.Y && other.Min.Y < this.Max.Y
}

func (r *Rect) FromImageRect(from image.Rectangle) {
	min := from.Min
	r.Min = Point{float64(min.X), float64(min.Y)}
	max := from.Max
	r.Max = Point{float64(max.X), float64(max.Y)}
}
