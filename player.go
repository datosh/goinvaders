package spaceinvaders

import (
	"github.com/hajimehoshi/ebiten"
)

type Player struct {
	*Sprite
}

func NewPlayer() *Player {
	player := &Player{
		Sprite: NewSprite("/canon.png"),
	}
	player.speed = 4
	player.x = 300
	player.y = 440

	return player
}

func (p *Player) Update(screen *ebiten.Image) error {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.MoveRelative(-1*p.speed, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.MoveRelative(1*p.speed, 0)
	}

	return nil
}
