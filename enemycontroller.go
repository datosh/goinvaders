package spaceinvaders

import (
	"spaceinvaders/vec2"

	"github.com/hajimehoshi/ebiten"
)

type EnemyController struct {
	Enemies      []*Enemy
	WalkableArea *Entity
}

func NewEnemyController() *EnemyController {
	ec := &EnemyController{
		WalkableArea: NewEntity(),
	}

	ec.AddEnemy(NewEnemy(&vec2.T{20, 20}, NewEnemy1Animation()))
	ec.AddEnemy(NewEnemy(&vec2.T{120, 20}, NewEnemy2Animation()))
	ec.AddEnemy(NewEnemy(&vec2.T{220, 20}, NewEnemy1Animation()))

	return ec
}

func (ec *EnemyController) AddEnemy(enemy *Enemy) {
	ec.Enemies = append(ec.Enemies, enemy)
}

func (ec *EnemyController) CollideWith(projectile *Projectile) {
	for _, enemy := range ec.Enemies {
		if enemy.Hitbox().Intersects(projectile.Hitbox()) {
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
	}
}

func (ec *EnemyController) RemoveDead() {
	ec.Enemies = Filter(ec.Enemies, isAlive).([]*Enemy)
}
