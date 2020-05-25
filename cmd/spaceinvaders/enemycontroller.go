package main

import (
	"engine"
	"engine/vec2"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type EnemyController struct {
	*engine.Entity
	Enemies         []*Enemy
	moveTimer       time.Time
	moveEach        time.Duration
	moveDistance    float64
	moveRight       bool
	changeDirection bool
	score           *Score
}

func NewEnemyController(score *Score) *EnemyController {
	ec := &EnemyController{
		Entity:          engine.NewEntity(),
		moveTimer:       time.Now(),
		moveEach:        time.Millisecond * 500,
		moveDistance:    20.0,
		moveRight:       true,
		changeDirection: false,
		score:           score,
	}
	ec.HitboxSize = &vec2.T{X: 570, Y: 400}
	ec.HitboxOffset = &vec2.T{X: 25, Y: 25}
	ec.Debug = false

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
			ec.score.Add(10)
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

	numEnemiesBefore := len(ec.Enemies)
	ec.RemoveDead()
	numEnemiesDied := numEnemiesBefore - len(ec.Enemies)
	ec.score.Add(100 * numEnemiesDied)
	return nil
}

func (ec *EnemyController) Draw(screen *ebiten.Image) {
	ec.Entity.Draw(screen)
	for _, enemy := range ec.Enemies {
		enemy.Draw(screen)
	}
}

func (ec *EnemyController) RemoveDead() {
	ec.Enemies = engine.Filter(ec.Enemies, isAlive).([]*Enemy)
}

type Killable interface {
	Dead() bool
}

func isAlive(elem interface{}) bool {
	entity := elem.(Killable)
	return !entity.Dead()
}
