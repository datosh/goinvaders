package engine

import (
	"image/color"

	"engine/vec2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Text struct {
	BaseNode
	Text     string
	Font     font.Face
	Position *vec2.T
}

func NewText(name, text string, font font.Face, position *vec2.T) *Text {
	return &Text{
		BaseNode: *NewNode(name),
		Text:     text,
		Font:     font,
		Position: position,
	}
}

func (t *Text) Draw(screen *ebiten.Image) {
	t.BaseNode.Draw(screen)
	text.Draw(
		screen, t.Text, t.Font,
		int(t.Position.X), int(t.Position.Y),
		color.White,
	)
}
