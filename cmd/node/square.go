package main

import (
	"engine"
	"engine/vec2"
)

type Square struct {
	engine.BaseNode
}

func NewSquare(game *engine.Game) *Square {
	square := &Square{
		BaseNode: *engine.NewNode("Square"),
	}

	globalPosition := engine.NewTransform("Position")
	globalPosition.Position = vec2.T{X: 300, Y: 100}
	square.AddChild(globalPosition)

	texture := engine.NewTexture(
		"SquareTexture",
		game.AssetLoader.LoadImage("assets/img/square.png"),
		globalPosition,
	)
	square.AddChild(texture)

	collider := engine.NewAABBCollider("SquareCollider", &globalPosition.Position, &vec2.T{X: 64, Y: 64})
	square.AddChild(collider)

	return square
}
