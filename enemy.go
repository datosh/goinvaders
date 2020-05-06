package spaceinvaders

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

type Enemy struct {
	img   *ebiten.Image
	x, y  float64
	scale float64
	speed float64
}

func NewEnemy(x, y float64, variant int) *Enemy {
	enemy := &Enemy{
		x:     x,
		y:     y,
		scale: 1,
		speed: 2,
	}
	enemy.img = LoadImage(fmt.Sprintf("/sprite%d.png", variant))
	return enemy
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(e.scale, e.scale)
	options.GeoM.Translate(e.x, e.y)
	screen.DrawImage(e.img, options)
}

func (e *Enemy) MoveRelative(x, y float64) {
	e.x += x
	e.y += y
}
