package engine

import (
	"engine/vec2"

	"github.com/hajimehoshi/ebiten/v2"
)

type Transform struct {
	BaseNode
	Position vec2.T
	Rotation float64
	Scale    vec2.T
}

func NewTransform(name string) *Transform {
	return &Transform{
		BaseNode: *NewNode(name),
		Position: vec2.T{X: 0, Y: 0},
		Rotation: 0,
		Scale:    vec2.T{X: 1, Y: 1},
	}
}

func (t *Transform) Matrix() ebiten.GeoM {
	geoM := ebiten.GeoM{}
	geoM.Scale(t.Scale.X, t.Scale.Y)
	geoM.Rotate(t.Rotation)
	geoM.Translate(t.Position.X, t.Position.Y)
	return geoM
}
