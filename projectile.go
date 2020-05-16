package spaceinvaders

import (
	"spaceinvaders/vec2"

	"github.com/hajimehoshi/ebiten"
)

type Projectile struct {
	*Sprite
}

func NewProjectile(x, y float64) *Projectile {
	projectile := &Projectile{
		Sprite: NewSprite(),
	}
	projectile.LoadImage(
		"/img/spritemap.png",
		TranslateBounds(vec2.Vec2I{64, 48}, vec2.Vec2I{2, 2}),
	)
	projectile.x = x
	projectile.y = y
	projectile.speed = 5.0
	projectile.scale = 0.25
	return projectile
}

func (p *Projectile) Update(screen *ebiten.Image) error {
	p.y -= 1 * p.speed

	return nil
}
