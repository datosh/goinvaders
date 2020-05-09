package spaceinvaders

import (
	"spaceinvaders/vec2"

	"github.com/hajimehoshi/ebiten"
)

type Player struct {
	*Sprite
}

func NewPlayer() *Player {
	player := &Player{
		Sprite: NewSprite(),
	}
	player.LoadImage(
		"/spritemap.png",
		TranslateBounds(vec2.Vec2I{64, 48}, vec2.Vec2I{2, 3}),
	)
	player.speed = 4
	player.x = 255
	player.y = 420
	player.scale = 1.2

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
