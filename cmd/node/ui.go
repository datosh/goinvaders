package main

import (
	"fmt"

	"engine"
	"engine/vec2"
)

type UI struct {
	engine.BaseNode
	healthText *engine.Text
}

func NewUI(game *engine.Game) *UI {
	ui := &UI{
		BaseNode: *engine.NewNode("UI"),
	}

	ui.healthText = engine.NewText(
		"PlayerHealth", "Health: 100",
		game.AssetLoader.LoadFont("assets/ttf/Orbitron.ttf", 14),
		vec2.New(10, 15))
	ui.AddChild(ui.healthText)

	return ui
}

// UpdateHealth updates the health display
func (u *UI) UpdateHealth(health int) {
	u.healthText.Text = fmt.Sprintf("Health: %d", health)
}
