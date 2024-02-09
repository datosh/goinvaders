package main

import (
	"math/rand"
	"time"

	"engine"
	"engine/vec2"

	"github.com/hajimehoshi/ebiten/v2"
)

type EnemyController struct {
	*engine.Entity
	Enemies                 []*Enemy
	fireTimer               time.Time
	fireEach                time.Duration
	projectiles             []*Projectile
	fireProjectileDirection *vec2.T
	moveTimer               time.Time
	moveEach                time.Duration
	moveDistance            float64
	moveRight               bool
	changeDirection         bool
	score                   *Score
}

func NewEnemyController(score *Score) *EnemyController {
	ec := &EnemyController{
		Entity:                  engine.NewEntity(),
		fireTimer:               time.Now(),
		fireEach:                time.Millisecond * 1200,
		fireProjectileDirection: vec2.UY().Mul(3),
		moveTimer:               time.Now(),
		moveEach:                time.Millisecond * 500,
		moveDistance:            20.0,
		moveRight:               true,
		changeDirection:         false,
		score:                   score,
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

func (ec *EnemyController) Update() error {
	ec.Entity.Update()

	for _, enemy := range ec.Enemies {
		enemy.Update()
	}

	for _, projectile := range ec.projectiles {
		projectile.Update()
	}

	// Do we make the nex step?
	if ec.moveTimer.Add(ec.moveEach).Before(time.Now()) {
		ec.moveTimer = time.Now()

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
	}

	if ec.fireTimer.Add(ec.fireEach).Before(time.Now()) {
		ec.fireTimer = time.Now()

		if len(ec.Enemies) > 0 {
			firingEnemy := ec.Enemies[rand.Intn(len(ec.Enemies))]
			ec.projectiles = append(
				ec.projectiles,
				NewProjectile(
					firingEnemy.Position.Added(vec2.New(30, 45)),
					ec.fireProjectileDirection,
				),
			)
		}
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
	for _, projectile := range ec.projectiles {
		projectile.Draw(screen)
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
