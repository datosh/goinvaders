package spaceinvaders

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

func DrawAABB(img *ebiten.Image, bounds Rect, clr color.Color) {
	rectImg, _ := ebiten.NewImage(
		int(bounds.w)+1,
		int(bounds.h)+1,
		ebiten.FilterDefault,
	)
	for x := 0; x < int(bounds.w); x++ {
		rectImg.Set(x, 0, clr)
		rectImg.Set(x, int(bounds.h), clr)
	}
	for y := 0; y < int(bounds.h); y++ {
		rectImg.Set(0, y, clr)
		rectImg.Set(int(bounds.w), y, clr)
	}
	img.DrawImage(rectImg, &ebiten.DrawImageOptions{
		GeoM: ebiten.TranslateGeo(bounds.x, bounds.y),
	})
}
