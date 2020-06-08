package main

import (
	"engine"
	"engine/vec2"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	starAnimationDuration = time.Millisecond * 250
)

type Star struct {
	*engine.Entity
	animation *engine.Animation
	game      ebiten.Game
}

func NewStarAnimation() *engine.Animation {
	return engine.NewAnimation(
		assetLoader.LoadImage("/img/Stern.png"),
		vec2.I{32, 32},
		[]vec2.I{
			{0, 0}, {1, 0}, {2, 0},
		},
		[]time.Duration{
			starAnimationDuration, starAnimationDuration, starAnimationDuration,
		},
	)
}

func NewStar(animation *engine.Animation, game ebiten.Game) *Star {
	star := &Star{
		Entity: engine.NewEntity(),
		game:   game,
	}
	star.animation = animation
	star.Image = star.animation.CurrentImage()
	star.randomLocation()

	go func() {
		for range time.Tick(starAnimationDuration * 5) {
			star.randomLocation()
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

func (s *Star) randomLocation() {
	windowWidth, windowHeight := s.game.Layout(0, 0)
	s.Position = vec2.NewI(
		rand.Intn(windowWidth),
		rand.Intn(windowHeight),
	).AsT()
}
