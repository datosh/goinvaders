package spaceinvaders

import "github.com/hajimehoshi/ebiten"

type Player struct {
	img   *ebiten.Image
	x, y  float64
	scale float64
	speed float64
}

func NewPlayer() *Player {
	player := &Player{
		x:     300,
		y:     440,
		scale: 1,
		speed: 2,
	}
	player.img = LoadImage("/canon.png")

	return player
}

func (p *Player) MoveRelative(x, y float64) {
	p.x += x
	p.y += y
}

func (p *Player) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(p.scale, p.scale)
	options.GeoM.Translate(p.x, p.y)
	screen.DrawImage(p.img, options)
}
