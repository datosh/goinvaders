package engine

import (
	"fmt"

	"engine/vec2"

	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	viewPort    *vec2.T
	position    *vec2.T
	zoom        *vec2.T
	rotation    float64
	worldMatrix ebiten.GeoM
}

func NewCamera(viewPort *vec2.T) *Camera {
	cam := &Camera{
		viewPort: viewPort,
		position: vec2.New(0, 0),
		zoom:     vec2.New(1, 1),
		rotation: 0,
	}
	cam.updateMatrix()
	return cam
}

func (c *Camera) String() string {
	return fmt.Sprintf(
		"T: %.2f, R: %.2f, S: %.2f, ViewPort: %.2f",
		c.position, c.rotation, c.zoom, c.viewPort,
	)
}

func (c *Camera) viewportCenter() *vec2.T {
	return c.viewPort.Muled(0.5)
}

func (c *Camera) updateMatrix() {
	c.worldMatrix.Reset()
	// First, move to center of screen, so our focal point is there,
	// not in top left corner. Then it is Translate * Rotate * Scale,
	// and move back to top left center
	// cf. https://gamedev.stackexchange.com/a/16721
	c.worldMatrix.Translate(c.position.Inverted().Coords())
	c.worldMatrix.Translate(c.viewportCenter().Inverted().Coords())
	c.worldMatrix.Scale(c.zoom.Coords())
	c.worldMatrix.Rotate(c.rotation)
	c.worldMatrix.Translate(c.viewportCenter().Coords())
}

func (c *Camera) Render(world, screen *ebiten.Image) {
	screen.DrawImage(world, &ebiten.DrawImageOptions{
		GeoM: c.worldMatrix,
	})
}

func (c *Camera) ScreenToWorld(point *vec2.T) *vec2.T {
	inverseMatrix := c.worldMatrix
	inverseMatrix.Invert()
	return vec2.New(inverseMatrix.Apply(point.Coords()))
}

func (c *Camera) WorldToScreen(point *vec2.T) *vec2.T {
	return vec2.New(c.worldMatrix.Apply(point.Coords()))
}

func (c *Camera) Move(delta *vec2.T) {
	c.position.Add(delta)
	c.updateMatrix()
}

func (c *Camera) MoveTo(pos *vec2.T) {
	c.position = pos.Copy()
	c.updateMatrix()
}

func (c *Camera) FocusOn(pos *vec2.T) {
	c.position = pos.Copy().Sub(c.viewportCenter())
	c.updateMatrix()
}

func (c *Camera) Zoom(m float64) {
	c.zoom.Mul(m)
	c.updateMatrix()
}

func (c *Camera) Rotate(theta float64) {
	c.rotation += theta
	c.updateMatrix()
}

func (c *Camera) Reset() {
	c.rotation = 0
	c.zoom.Set(1, 1)
	c.updateMatrix()
}
