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
	*Entity
	animation *Animation
}

func NewStarAnimation() *Animation {
	return NewAnimation(
		LoadImage("/img/Stern.png"),
		vec2.I{32, 32},
		[]vec2.I{
			{0, 0}, {1, 0}, {2, 0},
		},
		[]time.Duration{
			starAnimationDuration, starAnimationDuration, starAnimationDuration,
		},
	)
}

func NewStar(animation *Animation) *Star {
	star := &Star{
		Entity: NewEntity(),
	}
	star.animation = animation
	star.Image = star.animation.CurrentImage()
	ChangeStarLocation(star)

	go func() {
		for range time.Tick(starAnimationDuration * 5) {
			ChangeStarLocation(star)
		}
	}()

	return star
}

func (s *Star) Update(screen *ebiten.Image) error {
	s.Entity.Update(screen)
	s.animation.Update(screen)
	s.Image = s.animation.CurrentImage()
	return nil
}

func ChangeStarLocation(star *Star) {
	windowWidth, windowHeight := ebiten.WindowSize()
	star.Position = vec2.NewI(
		rand.Intn(windowWidth),
		rand.Intn(windowHeight),
	).AsT()
}
