package main

import (
	"engine"
	"engine/vec2"
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
)

type Score struct {
	*engine.Entity
	current int
	face    font.Face
}

func NewScore() *Score {
	score := &Score{
		Entity:  engine.NewEntity(),
		current: 0,
		face:    engine.LoadFont("/ttf/Orbitron.ttf", 14),
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
	s.current += delta
}

func (s *Score) text() string {
	return fmt.Sprintf("Score: %d", s.current)
}
