package spaceinvaders

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type EnemyController struct {
	Enemies      []*Enemy
	WalkableArea *Sprite
}

func NewEnemyController() *EnemyController {
	ec := &EnemyController{
		WalkableArea: NewSprite(),
	}

	DrawAABB(ec.WalkableArea.img, Rect{
		x: 20, y: 20, w: 200, h: 200,
	}, color.RGBA{0, 255, 0, 255})

	ec.AddEnemy(NewEnemy(20, 20, NewEnemy1Animation()))
	ec.AddEnemy(NewEnemy(120, 20, NewEnemy2Animation()))
	ec.AddEnemy(NewEnemy(220, 20, NewEnemy1Animation()))

	return ec
}

func (ec *EnemyController) AddEnemy(enemy *Enemy) {
	ec.Enemies = append(ec.Enemies, enemy)
}

func (ec *EnemyController) CollideWith(projectile *Projectile) {
	for _, enemy := range ec.Enemies {
		if DoCollide(enemy.Bounds(), projectile.Bounds()) {
			enemy.Hit()
			projectile.Die()
		}
	}
}

func (ec *EnemyController) Update(screen *ebiten.Image) error {
	for _, enemy := range ec.Enemies {
		enemy.Update(screen)
	}

	ec.RemoveDead()
	return nil
}

func (ec *EnemyController) Draw(screen *ebiten.Image) {
	for _, enemy := range ec.Enemies {
		enemy.Draw(screen)
		DrawAABB(screen, enemy.Bounds(), color.RGBA{0, 255, 0, 255})
	}
}

func (ec *EnemyController) RemoveDead() {
	ec.Enemies = Filter(ec.Enemies, isAlive).([]*Enemy)
}
