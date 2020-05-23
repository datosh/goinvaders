package spaceinvaders

import (
	"spaceinvaders/vec2"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
)

type Enemy struct {
	*Sprite
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
		vec2.PointI{X: 64, Y: 48},
		[]vec2.PointI{{X: 0, Y: 0}, {X: 1, Y: 0}},
		[]time.Duration{time.Millisecond * 750, time.Millisecond * 750},
	)
}

func NewEnemy2Animation() *Animation {
	return NewAnimation(
		LoadImage("/img/spritemap.png"),
		vec2.PointI{X: 64, Y: 48},
		[]vec2.PointI{{X: 2, Y: 0}, {X: 3, Y: 0}},
		[]time.Duration{time.Millisecond * 750, time.Millisecond * 750},
	)
}

func NewEnemy(position vec2.Point, animation *Animation) *Enemy {
	enemy := &Enemy{
		Sprite:       NewSprite(),
		moveTimer:    time.Now(),
		moveEach:     time.Millisecond * 750,
		moveDistance: 20.0,
		hitAudio:     LoadAudioPlayer("/audio/au.mp3"),
		hitPoints:    3,
	}
	enemy.animation = animation
	enemy.image = enemy.animation.CurrentImage()
	enemy.position = position
	return enemy
}

func (e *Enemy) Update(screen *ebiten.Image) error {
	if e.moveTimer.Add(e.moveEach).Before(time.Now()) {
		e.MoveRelativeX(e.moveDistance)
		e.moveTimer = time.Now()
	}
	e.animation.Update(screen)
	e.image = e.animation.CurrentImage()
	return nil
}

func (e *Enemy) Die() {
	e.alive = false
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
