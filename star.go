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

	thisAnimationDuration := time.Duration(starAnimationDuration + time.Duration(rand.Float64())*time.Millisecond*100)
	return NewAnimation(
		LoadImage("/img/Stern.png"),
		vec2.I{32, 32},
		[]vec2.I{
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
		Entity: NewEntity(),
	}
	star.animation = animation
	star.Image = star.animation.CurrentImage()
	ChangeStarLocation(star)

	go func() {
		for range time.Tick((starAnimationDuration * 5) + time.Duration(rand.Float64())*time.Millisecond*1520) {
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
	time.Sleep(time.Duration(int(rand.Float64()*20)) * time.Millisecond)
	star.Position = &vec2.T{
		X: rand.Float64() * 640,
		Y: rand.Float64() * 480,
	}
}
