package main

import (
	"embed"

	"engine"
)

//go:embed assets
var assets embed.FS

func main() {
	game := engine.NewGame("Testing Node Engine", assets)

	player := NewPlayer(game)
	game.Scene.AddNode(player)

	ui := NewUI(game)
	game.Scene.AddNode(ui)

	player.OnHealthChanged.Connect(ui.UpdateHealth)

	enemy := NewEnemy(game)
	game.Scene.AddNode(enemy)

	square := NewSquare(game)
	game.Scene.AddNode(square)

	game.Run()
}
