package spaceinvaders

import (
	"spaceinvaders/vec2"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type EnemyController struct {
	*Entity
	Enemies         []*Enemy
	moveTimer       time.Time
	moveEach        time.Duration
	moveDistance    float64
	moveRight       bool
	changeDirection bool
}

func NewEnemyController() *EnemyController {
	ec := &EnemyController{
		Entity:          NewEntity(),
		moveTimer:       time.Now(),
		moveEach:        time.Millisecond * 500,
		moveDistance:    20.0,
		moveRight:       true,
		changeDirection: false,
	}
	ec.HitboxSize = &vec2.T{570, 400}
	ec.HitboxOffset = &vec2.T{25, 25}
	ec.Debug = false

	ec.AddEnemy(NewEnemy(&vec2.T{020, 30}, NewEnemy1Animation()))
	ec.AddEnemy(NewEnemy(&vec2.T{120, 30}, NewEnemy1Animation()))
	ec.AddEnemy(NewEnemy(&vec2.T{220, 30}, NewEnemy1Animation()))
	ec.AddEnemy(NewEnemy(&vec2.T{320, 30}, NewEnemy1Animation()))
	ec.AddEnemy(NewEnemy(&vec2.T{420, 30}, NewEnemy1Animation()))

	ec.AddEnemy(NewEnemy(&vec2.T{020, 150}, NewEnemy2Animation()))
	ec.AddEnemy(NewEnemy(&vec2.T{120, 150}, NewEnemy2Animation()))
	ec.AddEnemy(NewEnemy(&vec2.T{220, 150}, NewEnemy2Animation()))
	ec.AddEnemy(NewEnemy(&vec2.T{320, 150}, NewEnemy2Animation()))
	ec.AddEnemy(NewEnemy(&vec2.T{420, 150}, NewEnemy2Animation()))

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

	// Do we make the nex step?
	if ec.moveTimer.Add(ec.moveEach).Before(time.Now()) {

		moveDirection := vec2.UX().Mul(ec.moveDistance)
		if !ec.moveRight {
			moveDirection.Invert()
		}
		if ec.changeDirection {
			moveDirection = vec2.UY().Mul(ec.moveDistance)
			ec.changeDirection = false
			ec.moveRight = !ec.moveRight
		}

		for _, enemy := range ec.Enemies {
			enemy.Position.Add(moveDirection)
			if !enemy.Hitbox().Inside(ec.Hitbox()) && moveDirection.Y == 0.0 {
				ec.changeDirection = true
			}
		}
		ec.moveTimer = time.Now()
	}

	ec.RemoveDead()
	return nil
}

func (ec *EnemyController) Draw(screen *ebiten.Image) {
	ec.Entity.Draw(screen)
	for _, enemy := range ec.Enemies {
		enemy.Draw(screen)
	}
}

func (ec *EnemyController) RemoveDead() {
	ec.Enemies = Filter(ec.Enemies, isAlive).([]*Enemy)
}
