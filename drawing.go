package engine

import (
	"image/color"

	"engine/vec2"

	"github.com/hajimehoshi/ebiten/v2"
)

// DrawRect draws a one pixel wide border of the rectangle specified by
// `bounds` on `img`.
func DrawRect(img *ebiten.Image, bounds vec2.Rect, clr color.Color) {
	rectImg := ebiten.NewImage(
		int(bounds.Width()),
		int(bounds.Height()),
	)
	for x := 0; x < int(bounds.Width()); x++ {
		rectImg.Set(x, 0, clr)
		rectImg.Set(x, int(bounds.Height())-1, clr)
	}
	for y := 0; y < int(bounds.Height()); y++ {
		rectImg.Set(0, y, clr)
		rectImg.Set(int(bounds.Width())-1, y, clr)
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(bounds.X(), bounds.Y())
	img.DrawImage(rectImg, op)
}
