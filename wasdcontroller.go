package engine

import (
	"engine/vec2"

	"github.com/hajimehoshi/ebiten/v2"
)

type WASDController struct {
	BaseNode
	controlledPosition *vec2.T

	Speed float64
	WKey  ebiten.Key
	AKey  ebiten.Key
	SKey  ebiten.Key
	DKey  ebiten.Key
}

func NewWASDController(speed float64, controlledPosition *vec2.T) *WASDController {
	return &WASDController{
		BaseNode:           *NewNode("WASDController"),
		controlledPosition: controlledPosition,
		Speed:              speed,
		WKey:               ebiten.KeyW,
		AKey:               ebiten.KeyA,
		SKey:               ebiten.KeyS,
		DKey:               ebiten.KeyD,
	}
}

func (c *WASDController) Update() error {
	direction := vec2.T{}
	if ebiten.IsKeyPressed(c.AKey) {
		direction.Add(vec2.UX().Invert())
	}
	if ebiten.IsKeyPressed(c.DKey) {
		direction.Add(vec2.UX())
	}
	if ebiten.IsKeyPressed(c.WKey) {
		direction.Add(vec2.UY().Invert())
	}
	if ebiten.IsKeyPressed(c.SKey) {
		direction.Add(vec2.UY())
	}

	direction.Normalize()
	c.controlledPosition.Add(direction.Mul(c.Speed))

	return nil
}
