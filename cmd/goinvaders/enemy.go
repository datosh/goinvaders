package main

import (
	"log"
	"time"

	"engine"
	"engine/vec2"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

type EnemyType int

const (
	EnemyOne EnemyType = iota
	EnemyTwo
)

type Enemy struct {
	*engine.Entity
	animation *engine.Animation
	hitAudio  *audio.Player
	hitPoints int
}

func newEnemy1Animation() *engine.Animation {
	return engine.NewAnimation(
		assetLoader.LoadImage("assets/img/spritemap.png"),
		vec2.I{X: 64, Y: 48},
		[]vec2.I{{X: 0, Y: 0}, {X: 1, Y: 0}},
		engine.UniformDuration(time.Millisecond*500, 2),
	)
}

func newEnemy2Animation() *engine.Animation {
	return engine.NewAnimation(
		assetLoader.LoadImage("assets/img/spritemap.png"),
		vec2.I{X: 64, Y: 48},
		[]vec2.I{{X: 3, Y: 0}, {X: 2, Y: 0}},
		engine.UniformDuration(time.Millisecond*500, 2),
	)
}

func NewEnemy(position *vec2.T, enemyType EnemyType) *Enemy {
	enemy := &Enemy{
		Entity:    engine.NewEntity(),
		hitAudio:  assetLoader.LoadAudioPlayer("assets/audio/au.mp3"),
		hitPoints: 3,
	}

	switch enemyType {
	case EnemyOne:
		enemy.animation = newEnemy1Animation()
	case EnemyTwo:
		enemy.animation = newEnemy2Animation()
	default:
		log.Panicf("There is no EnemyType: %v", enemyType)
	}

	enemy.Image = enemy.animation.CurrentImage()
	enemy.Position = position
	enemy.HitboxSize = vec2.NewI(
		enemy.Image.Bounds().Size().X,
		enemy.Image.Bounds().Size().Y,
	).AsT()

	return enemy
}

func (e *Enemy) Update() error {
	e.Entity.Update()
	e.animation.Update()
	e.Image = e.animation.CurrentImage()
	return nil
}

func (e *Enemy) Hit() {
	e.hitPoints--
	e.hitAudio.Rewind()
	e.hitAudio.Play()
	if e.hitPoints == 0 {
		e.Die()
	}
}
