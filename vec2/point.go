package vec2

import "math"

type Point struct {
	X, Y float64
}

type PointI struct {
	X, Y int
}

func (self *Point) Add(other Point) {
	self.X += other.X
	self.Y += other.Y
}

func (self *Point) Sub(other Point) {
	self.X -= other.X
	self.Y -= other.Y
}

func (self *Point) Mul(other float64) {
	self.X *= other
	self.Y *= other
}

func (self Point) Length() float64 {
	return math.Sqrt(self.X*self.X + self.Y*self.Y)
}

func (self *Point) Normalize() {
	self.Mul(1 / self.Length())
}

func (self *Point) AsPointI() PointI {
	return PointI{int(self.X), int(self.Y)}
}

func (self *PointI) AsPoint() Point {
	return Point{X: float64(self.X), Y: float64(self.Y)}
}
