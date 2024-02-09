package main

import (
	"engine"
	"engine/vec2"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
)

type Score struct {
	*engine.Entity
	currentScore int
	currentLevel int
	face         font.Face
}

func NewScore() *Score {
	score := &Score{
		Entity:       engine.NewEntity(),
		currentScore: 0,
		currentLevel: 1,
		face:         assetLoader.LoadFont("assets/ttf/Orbitron.ttf", 14),
	}
	score.Position.Add(vec2.New(10, 15))

	return score
}

func (s *Score) Draw(screen *ebiten.Image) {
	text.Draw(
		screen, s.text(), s.face,
		s.Position.AsI().X, s.Position.AsI().Y,
		colornames.Yellow,
	)
}

func (s *Score) Add(delta int) {
	s.currentScore += delta
}

func (s *Score) SetLevel(level int) {
	s.currentLevel = level
}

func (s *Score) text() string {
	return fmt.Sprintf("Level: %d    Score: %d", s.currentLevel, s.currentScore)
}
