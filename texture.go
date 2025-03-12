package engine

import (
	"engine/vec2"

	"github.com/hajimehoshi/ebiten/v2"
)

type Texture struct {
	BaseNode
	Image           *ebiten.Image
	Offset          vec2.T
	Scale           vec2.T
	Rotation        float64
	GlobalTransform *Transform
}

func NewTexture(name string, image *ebiten.Image, globalTransform *Transform) *Texture {
	return &Texture{
		BaseNode:        *NewNode(name),
		Image:           image,
		Offset:          vec2.T{X: 0, Y: 0},
		Scale:           vec2.T{X: 1, Y: 1},
		Rotation:        0,
		GlobalTransform: globalTransform,
	}
}

func (t *Texture) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	// Local translation
	op.GeoM.Scale(t.Scale.X, t.Scale.Y)
	op.GeoM.Rotate(t.Rotation)
	op.GeoM.Translate(t.Offset.X, t.Offset.Y)

	op.GeoM.Concat(t.GlobalTransform.Matrix())

	screen.DrawImage(t.Image, op)
}
