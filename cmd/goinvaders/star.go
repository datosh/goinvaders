package main

import (
	"math/rand"
	"time"

	"engine"
	"engine/vec2"

	"github.com/hajimehoshi/ebiten/v2"
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
		assetLoader.LoadImage("assets/img/Stern.png"),
		vec2.I{X: 32, Y: 32},
		[]vec2.I{
			{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0},
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

func (s *Star) Update() error {
	s.Entity.Update()
	s.animation.Update()
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
