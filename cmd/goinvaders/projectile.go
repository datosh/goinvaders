package main

import (
	"engine"
	"engine/vec2"

	"github.com/hajimehoshi/ebiten/v2"
)

type Projectile struct {
	*engine.Entity
	Direction *vec2.T
}

func NewProjectile(position, direction *vec2.T) *Projectile {
	projectile := &Projectile{
		Entity:    engine.NewEntity(),
		Direction: direction,
	}
	projectile.Image = assetLoader.LoadSubImage(
		"assets/img/spritemap.png",
		engine.CoordinatesToBounds(vec2.I{X: 64, Y: 48}, vec2.I{X: 2, Y: 2}),
	)
	projectile.Position = position
	projectile.ImageScale = 0.25
	projectile.HitboxSize = vec2.NewI(projectile.Image.Size()).AsT().
		Mul(projectile.ImageScale)
	return projectile
}

func (p *Projectile) Update() error {
	p.Entity.Update()
	p.Position.Add(p.Direction)
	return nil
}

func (p *Projectile) Draw(screen *ebiten.Image) {
	p.Entity.Draw(screen)
}
