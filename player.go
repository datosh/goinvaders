package spaceinvaders

import (
	"spaceinvaders/vec2"

	"github.com/hajimehoshi/ebiten"
)

type Player struct {
	*Sprite
	speed float64
}

func NewPlayer() *Player {
	player := &Player{
		Sprite: NewSprite(),
	}
	player.SetImage(LoadSubImage(
		"/img/spritemap.png",
		CoordinatesToBounds(vec2.I{64, 48}, vec2.I{2, 3}),
	))
	player.speed = 4
	player.SetPosition(&vec2.T{255, 420})
	player.imageScale = 1.2

	return player
}

func (p *Player) Update(screen *ebiten.Image) error {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.MoveRelativeX(-1 * p.speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.MoveRelativeX(1 * p.speed)
	}

	return nil
}
