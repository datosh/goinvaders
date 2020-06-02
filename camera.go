package engine

import (
	"engine/vec2"

	"github.com/hajimehoshi/ebiten"
)

type Camera struct {
	position *vec2.T
}

func NewCamera(x, y float64) *Camera {
	return &Camera{
		position: vec2.New(x, y),
	}
}

func (c *Camera) View(world, screen *ebiten.Image) error {
	geoM := ebiten.GeoM{}
	geoM.Translate(c.position.Inverted().Coords())
	return screen.DrawImage(world, &ebiten.DrawImageOptions{
		GeoM: geoM,
	})
}

func (c *Camera) ScreenToWorld(point vec2.T) *vec2.T {
	return point.Sub(c.position)
}

func (c *Camera) Move(delta *vec2.T) {
	c.position.Add(delta)
}
