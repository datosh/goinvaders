package main

import (
	"engine"
	"engine/vec2"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	*engine.Entity
}

func NewEnemy() *Enemy {
	enemy := &Enemy{Entity: engine.NewEntity()}

	image := assetLoader.LoadImage("assets/img/enemy.png")
	enemy.Image = image
	enemy.ImageScale = 0.5
	enemy.Position = &vec2.T{X: 900, Y: 700}
	enemy.HitboxSize = &vec2.T{X: 64, Y: 64}
	enemy.HitboxOffset = &vec2.T{X: 0, Y: 0}

	return enemy
}

func (e *Enemy) Update() error {
	return nil
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	e.Entity.Draw(screen)
}
