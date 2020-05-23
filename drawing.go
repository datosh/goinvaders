package spaceinvaders

import (
	"image/color"
	"spaceinvaders/vec2"

	"github.com/hajimehoshi/ebiten"
)

func DrawAABB(img *ebiten.Image, bounds vec2.Rect, clr color.Color) {
	rectImg, _ := ebiten.NewImage(
		int(bounds.Width())+1,
		int(bounds.Height())+1,
		ebiten.FilterDefault,
	)
	for x := 0; x < int(bounds.Width()); x++ {
		rectImg.Set(x, 0, clr)
		rectImg.Set(x, int(bounds.Height()), clr)
	}
	for y := 0; y < int(bounds.Height()); y++ {
		rectImg.Set(0, y, clr)
		rectImg.Set(int(bounds.Width()), y, clr)
	}
	img.DrawImage(rectImg, &ebiten.DrawImageOptions{
		GeoM: ebiten.TranslateGeo(bounds.X(), bounds.Y()),
	})
}
