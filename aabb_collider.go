package engine

import (
	"image/color"

	"engine/vec2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type AABBCollider struct {
	BaseNode
	Position    *vec2.T
	Size        *vec2.T
	OnCollision *Signal[Node]
	DebugMode   bool
}

func NewAABBCollider(name string, position *vec2.T, size *vec2.T) *AABBCollider {
	collider := &AABBCollider{
		BaseNode:    *NewNode(name),
		Position:    position,
		Size:        size,
		OnCollision: NewSignal[Node](),
	}
	collider.AddTag("Collider")
	return collider
}

func (c *AABBCollider) Update() error {
	if err := c.BaseNode.Update(); err != nil {
		return err
	}

	if ebiten.IsKeyPressed(ebiten.KeyF10) {
		c.DebugMode = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyControl) && ebiten.IsKeyPressed(ebiten.KeyF10) {
		c.DebugMode = false
	}

	return nil
}

func (c *AABBCollider) Draw(screen *ebiten.Image) {
	c.BaseNode.Draw(screen)

	if c.DebugMode {
		vector.StrokeRect(
			screen,
			float32(c.Position.X),
			float32(c.Position.Y),
			float32(c.Size.X),
			float32(c.Size.Y),
			1,
			color.White,
			false,
		)
	}
}

func (c *AABBCollider) CollidesWithAABB(other *AABBCollider) bool {
	return c.Position.X < other.Position.X+other.Size.X &&
		c.Position.X+c.Size.X > other.Position.X &&
		c.Position.Y < other.Position.Y+other.Size.Y &&
		c.Position.Y+c.Size.Y > other.Position.Y
}

func (c *AABBCollider) CollidesWithCircle(other *CircleCollider) bool {
	// Check if the circle is inside the AABB
	if other.Position.X >= c.Position.X &&
		other.Position.X+other.Radius <= c.Position.X+c.Size.X &&
		other.Position.Y >= c.Position.Y &&
		other.Position.Y+other.Radius <= c.Position.Y+c.Size.Y {
		return true
	}

	// Check if the AABB is inside the circle
	if c.Position.X >= other.Position.X &&
		c.Position.X+c.Size.X <= other.Position.X+other.Radius &&
		c.Position.Y >= other.Position.Y &&
		c.Position.Y+c.Size.Y <= other.Position.Y+other.Radius {
		return true
	}

	// Check if the AABB is outside the circle
	if c.Position.X > other.Position.X+other.Radius ||
		c.Position.X+c.Size.X < other.Position.X ||
		c.Position.Y > other.Position.Y+other.Radius ||
		c.Position.Y+c.Size.Y < other.Position.Y {
		return false
	}

	return true
}
