package spaceinvaders

import "github.com/hajimehoshi/ebiten"

type Sprite struct {
	img   *ebiten.Image
	x, y  float64
	scale float64
	speed float64
}

func NewSprite(path string) *Sprite {
	sprite := &Sprite{
		img:   LoadImage(path),
		x:     0,
		y:     0,
		scale: 1,
		speed: 1,
	}
	return sprite
}

func (s *Sprite) MoveRelative(x, y float64) {
	s.x += x
	s.y += y
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(s.scale, s.scale)
	options.GeoM.Translate(s.x, s.y)
	screen.DrawImage(s.img, options)
}
