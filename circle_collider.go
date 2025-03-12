package engine

import (
	"image/color"
	"math"

	"engine/vec2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type CircleCollider struct {
	BaseNode
	Position    *vec2.T
	Radius      float64
	OnCollision *Signal[Node]
	DebugMode   bool
}

func NewCircleCollider(name string, position *vec2.T, radius float64) *CircleCollider {
	collider := &CircleCollider{
		BaseNode:    *NewNode(name),
		Position:    position,
		Radius:      radius,
		OnCollision: NewSignal[Node](),
	}
	collider.AddTag("Collider")
	return collider
}

func (c *CircleCollider) Update() error {
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

func (c *CircleCollider) Draw(screen *ebiten.Image) {
	c.BaseNode.Draw(screen)

	if c.DebugMode {
		vector.StrokeCircle(
			screen,
			float32(c.Position.X+c.Radius),
			float32(c.Position.Y+c.Radius),
			float32(c.Radius),
			1,
			color.White,
			false,
		)
	}
}

func (c *CircleCollider) CollidesWithCircle(other *CircleCollider) bool {
	// Calculate distance between centers
	dx := c.Position.X + c.Radius - (other.Position.X + other.Radius)
	dy := c.Position.Y + c.Radius - (other.Position.Y + other.Radius)
	distance := math.Sqrt(dx*dx + dy*dy)

	// If distance is less than sum of radii, they are colliding
	return distance < (c.Radius + other.Radius)
}

func (c *CircleCollider) CollidesWithAABB(other *AABBCollider) bool {
	return other.CollidesWithCircle(c)
}
