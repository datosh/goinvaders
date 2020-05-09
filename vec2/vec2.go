package vec2

import "math"

type Vec2 struct {
	X, Y float64
}

type Vec2I struct {
	X, Y int
}

func (self *Vec2) Add(other Vec2) {
	self.X += other.X
	self.Y += other.Y
}

func (self *Vec2) Sub(other Vec2) {
	self.X -= other.X
	self.Y -= other.Y
}

func (self *Vec2) Mul(other float64) {
	self.X *= other
	self.Y *= other
}

func (self Vec2) Length() float64 {
	return math.Sqrt(self.X*self.X + self.Y*self.Y)
}

func (self *Vec2) Normalize() {
	self.Mul(1 / self.Length())
}

func (self *Vec2) AsVec2I() Vec2I {
	return Vec2I{int(self.X), int(self.Y)}
}
