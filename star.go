package spaceinvaders

import (
	"math/rand"
	"spaceinvaders/vec2"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	starAnimationDuration = time.Millisecond * 250
)

type Star struct {
	*Sprite
	animation *Animation
}

func NewStarAnimation(initialDelay time.Duration) *Animation {

	return NewAnimation(
		"/Stern.png",
		vec2.Vec2I{32, 32},
		[]vec2.Vec2I{
			{0, 0}, {1, 0}, {2, 0},
			{0, 1}, {1, 1},
		},
		[]time.Duration{
			starAnimationDuration + initialDelay, starAnimationDuration, starAnimationDuration,
			starAnimationDuration, starAnimationDuration,
		},
	)
}

func NewStar(x, y float64, animation *Animation) *Star {
	star := &Star{
		Sprite: NewSprite(),
	}
	star.animation = animation
	star.img = star.animation.CurrentImage()
	star.x = x
	star.y = y

	go func() {
		for range time.Tick((starAnimationDuration * 5) + time.Duration(int(rand.Float64()*520))*time.Millisecond) {
			ChangeStarLocation(star)
		}
	}()

	return star
}

func (e *Star) Update(screen *ebiten.Image) error {
	e.animation.Update(screen)
	e.img = e.animation.CurrentImage()
	return nil
}

func ChangeStarLocation(star *Star) {
	time.Sleep(time.Duration(int(rand.Float64()*20)) * time.Millisecond)
	star.x = rand.Float64() * 640
	star.y = rand.Float64() * 480
}
