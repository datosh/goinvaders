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

func NewStarAnimation() *Animation {

	thisAnimationDuration := time.Duration(starAnimationDuration + time.Duration(rand.Float64())*time.Millisecond*100)
	return NewAnimation(
		LoadImage("/img/Stern.png"),
		vec2.PointI{32, 32},
		[]vec2.PointI{
			{0, 0}, {1, 0}, {2, 0},
			{0, 1}, {1, 1},
		},
		[]time.Duration{
			thisAnimationDuration, thisAnimationDuration, thisAnimationDuration,
			thisAnimationDuration, thisAnimationDuration,
		},
	)
}

func NewStar(animation *Animation) *Star {
	star := &Star{
		Sprite: NewSprite(),
	}
	star.animation = animation
	star.img = star.animation.CurrentImage()
	ChangeStarLocation(star)

	go func() {
		for range time.Tick((starAnimationDuration * 5) + time.Duration(rand.Float64())*time.Millisecond*1520) {
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
