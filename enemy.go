package spaceinvaders

import (
	"spaceinvaders/vec2"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type Enemy struct {
	*Sprite
	moveTimer    time.Time
	moveEach     time.Duration
	moveDistance float64
	animation    *Animation
}

func NewEnemy1Animation() *Animation {
	return NewAnimation(
		LoadImage("/spritemap.png"),
		vec2.Vec2I{64, 48},
		[]vec2.Vec2I{{0, 0}, {1, 0}},
		[]time.Duration{time.Millisecond * 750, time.Millisecond * 750},
	)
}

func NewEnemy2Animation() *Animation {
	return NewAnimation(
		LoadImage("/spritemap.png"),
		vec2.Vec2I{64, 48},
		[]vec2.Vec2I{{2, 0}, {3, 0}},
		[]time.Duration{time.Millisecond * 750, time.Millisecond * 750},
	)
}

func NewEnemy(x, y float64, animation *Animation) *Enemy {
	enemy := &Enemy{
		Sprite:       NewSprite(),
		moveTimer:    time.Now(),
		moveEach:     time.Millisecond * 750,
		moveDistance: 20.0,
	}
	enemy.animation = animation
	enemy.img = enemy.animation.CurrentImage()
	enemy.x = x
	enemy.y = y
	return enemy
}

func (e *Enemy) Update(screen *ebiten.Image) error {
	if e.moveTimer.Add(e.moveEach).Before(time.Now()) {
		e.MoveRelative(e.moveDistance, 0)
		e.moveTimer = time.Now()
	}
	e.animation.Update(screen)
	e.img = e.animation.CurrentImage()
	return nil
}
