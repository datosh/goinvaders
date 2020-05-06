package spaceinvaders

import "github.com/hajimehoshi/ebiten"

type Projectile struct {
	img   *ebiten.Image
	x, y  float64
	scale float64
	speed float64
}

func NewProjectile(x, y float64) *Projectile {
	projectile := &Projectile{
		x:     x,
		y:     y,
		scale: 1,
		speed: 2,
	}
	projectile.img = LoadImage("/projectile.png")

	return projectile
}

func (p *Projectile) Update(screen *ebiten.Image) error {
	p.y -= 1 * p.speed

	return nil
}

func (p *Projectile) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(p.scale, p.scale)
	options.GeoM.Translate(p.x, p.y)
	screen.DrawImage(p.img, options)
}
