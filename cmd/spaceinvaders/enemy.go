package main

import (
	"engine"
	"engine/vec2"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
)

type Enemy struct {
	*engine.Entity
	animation *engine.Animation
	hitAudio  *audio.Player
	hitPoints int
}

func NewEnemy1Animation() *engine.Animation {
	return engine.NewAnimation(
		engine.LoadImage("/img/spritemap.png"),
		vec2.I{X: 64, Y: 48},
		[]vec2.I{{X: 0, Y: 0}, {X: 1, Y: 0}},
		[]time.Duration{time.Millisecond * 500, time.Millisecond * 500},
	)
}

func NewEnemy2Animation() *engine.Animation {
	return engine.NewAnimation(
		engine.LoadImage("/img/spritemap.png"),
		vec2.I{X: 64, Y: 48},
		[]vec2.I{{X: 3, Y: 0}, {X: 2, Y: 0}},
		[]time.Duration{time.Millisecond * 500, time.Millisecond * 500},
	)
}

func NewEnemy(position *vec2.T, animation *engine.Animation) *Enemy {
	enemy := &Enemy{
		Entity:    engine.NewEntity(),
		hitAudio:  engine.LoadAudioPlayer("/audio/au.mp3"),
		hitPoints: 3,
	}
	enemy.animation = animation
	enemy.Image = enemy.animation.CurrentImage()
	enemy.Position = position
	enemy.HitboxSize = vec2.NewI(enemy.Image.Size()).AsT()
	return enemy
}

func (e *Enemy) Update(screen *ebiten.Image) error {
	e.Entity.Update(screen)
	e.animation.Update(screen)
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