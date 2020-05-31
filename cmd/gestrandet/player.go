package main

import (
	"engine"
	"engine/vec2"
	"log"

	"github.com/hajimehoshi/ebiten"
)

type Player struct {
	*engine.Entity
	speed float64
}

func NewPlayer() *Player {
	player := &Player{
		Entity: engine.NewEntity(),
		speed:  4,
	}
	player.Image = assetLoader.LoadSubImage(
		"/img/Character Spritemap.png",
		engine.CoordinatesToBounds(vec2.I{40, 56}, vec2.I{0, 0}),
	)
	player.Position = &vec2.T{255, 420}
	player.ImageScale = 1.5
	player.HitboxSize = vec2.New(64, 48)

	return player
}

func (p *Player) Update(screen *ebiten.Image) error {
	p.Entity.Update(screen)

	direction := vec2.New(0, 0)
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		direction.Add(vec2.UX().Invert())
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		direction.Add(vec2.UX())
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		direction.Add(vec2.UY().Invert())
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		direction.Add(vec2.UY())
	}
	direction.Normalize()
	direction.Mul(p.speed)
	log.Printf("Move into %v", direction)
	p.Position.Add(direction)

	log.Println(p.Position)

	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.Entity.Draw(screen)
}
