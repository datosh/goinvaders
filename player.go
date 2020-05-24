package spaceinvaders

import (
	"spaceinvaders/vec2"

	"github.com/hajimehoshi/ebiten"
)

type Player struct {
	*Entity
	speed float64
}

func NewPlayer() *Player {
	player := &Player{
		Entity: NewEntity(),
	}
	player.Image = LoadSubImage(
		"/img/spritemap.png",
		CoordinatesToBounds(vec2.I{64, 48}, vec2.I{2, 3}),
	)
	player.speed = 4
	player.Position = &vec2.T{255, 420}
	player.ImageScale = 1.2

	return player
}

func (p *Player) Update(screen *ebiten.Image) error {
	p.Entity.Update(screen)
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.Position.Add(vec2.UX().Mul(p.speed).Invert())
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.Position.Add(vec2.UX().Mul(p.speed))
	}

	return nil
}
