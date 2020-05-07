package spaceinvaders

import "github.com/hajimehoshi/ebiten"

type Projectile struct {
	*Sprite
}

func NewProjectile(x, y float64) *Projectile {
	projectile := &Projectile{
		Sprite: NewSprite("/projectile.png"),
	}
	projectile.x = x
	projectile.y = y
	projectile.speed = 5.0
	return projectile
}

func (p *Projectile) Update(screen *ebiten.Image) error {
	p.y -= 1 * p.speed

	return nil
}
