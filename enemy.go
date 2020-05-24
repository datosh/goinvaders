package spaceinvaders

import (
	"spaceinvaders/vec2"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
)

type Enemy struct {
	*Entity
	moveTimer    time.Time
	moveEach     time.Duration
	moveDistance float64
	animation    *Animation
	hitAudio     *audio.Player
	hitPoints    int
}

func NewEnemy1Animation() *Animation {
	return NewAnimation(
		LoadImage("/img/spritemap.png"),
		vec2.I{X: 64, Y: 48},
		[]vec2.I{{X: 0, Y: 0}, {X: 1, Y: 0}},
		[]time.Duration{time.Millisecond * 750, time.Millisecond * 750},
	)
}

func NewEnemy2Animation() *Animation {
	return NewAnimation(
		LoadImage("/img/spritemap.png"),
		vec2.I{X: 64, Y: 48},
		[]vec2.I{{X: 2, Y: 0}, {X: 3, Y: 0}},
		[]time.Duration{time.Millisecond * 750, time.Millisecond * 750},
	)
}

func NewEnemy(position *vec2.T, animation *Animation) *Enemy {
	enemy := &Enemy{
		Entity:       NewEntity(),
		moveTimer:    time.Now(),
		moveEach:     time.Millisecond * 750,
		moveDistance: 20.0,
		hitAudio:     LoadAudioPlayer("/audio/au.mp3"),
		hitPoints:    3,
	}
	enemy.animation = animation
	enemy.Image = enemy.animation.CurrentImage()
	enemy.Position = position
	enemy.HitboxSize = vec2.NewI(enemy.Image.Size()).AsT()
	enemy.HitboxOffset = vec2.New(1, 1)
	return enemy
}

func (e *Enemy) Update(screen *ebiten.Image) error {
	e.Entity.Update(screen)
	if e.moveTimer.Add(e.moveEach).Before(time.Now()) {
		e.Position.Add(vec2.UX().Mul(e.moveDistance))
		e.moveTimer = time.Now()
	}
	e.animation.Update(screen)
	e.Image = e.animation.CurrentImage()
	return nil
}

func (e *Enemy) Die() {
	e.Alive = false
}

func (e *Enemy) Hit() {
	e.hitPoints--
	e.hitAudio.SetVolume(1)
	e.hitAudio.Rewind()
	e.hitAudio.Play()
	if e.hitPoints == 0 {
		e.Die()
	}
}
