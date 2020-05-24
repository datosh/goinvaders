package spaceinvaders

import (
	"spaceinvaders/vec2"

	"github.com/hajimehoshi/ebiten"
)

type Projectile struct {
	*Entity
	speed float64
}

func NewProjectile(position *vec2.T) *Projectile {
	projectile := &Projectile{
		Entity: NewEntity(),
	}
	projectile.Image = LoadSubImage(
		"/img/spritemap.png",
		CoordinatesToBounds(vec2.I{64, 48}, vec2.I{2, 2}),
	)
	projectile.Position = position
	projectile.speed = 5.0
	projectile.ImageScale = 0.25
	projectile.HitboxSize = vec2.NewI(projectile.Image.Size()).AsT().
		Mul(projectile.ImageScale)
	return projectile
}

func (p *Projectile) Update(screen *ebiten.Image) error {
	p.Entity.Update(screen)
	p.Position.Add(vec2.UY().Mul(p.speed).Invert())
	return nil
}

func (p *Projectile) Draw(screen *ebiten.Image) {
	p.Entity.Draw(screen)
}
