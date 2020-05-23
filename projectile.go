package spaceinvaders

import (
	"image/color"
	"spaceinvaders/vec2"

	"github.com/hajimehoshi/ebiten"
)

type Projectile struct {
	*Sprite
	speed float64
}

func NewProjectile(position vec2.Point) *Projectile {
	projectile := &Projectile{
		Sprite: NewSprite(),
	}
	projectile.LoadSubImage(
		"/img/spritemap.png",
		CoordinatesToBounds(vec2.PointI{64, 48}, vec2.PointI{2, 2}),
	)
	projectile.position = position
	projectile.speed = 5.0
	projectile.imageScale = 0.25
	return projectile
}

func (p *Projectile) Update(screen *ebiten.Image) error {
	p.MoveRelativeY(-1 * p.speed)
	return nil
}

func (p *Projectile) Draw(screen *ebiten.Image) {
	p.Sprite.Draw(screen)
	DrawAABB(screen, p.ImageBounds(), color.RGBA{0, 255, 0, 255})
}
