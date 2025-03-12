package main

import (
	"engine"
	"engine/vec2"
)

type Enemy struct {
	engine.BaseNode
}

func NewEnemy(game *engine.Game) *Enemy {
	enemy := &Enemy{
		BaseNode: *engine.NewNode("Enemy"),
	}

	globalPosition := engine.NewTransform("Position")
	globalPosition.Position = vec2.T{X: 100, Y: 200}
	enemy.AddChild(globalPosition)

	texture := engine.NewTexture(
		"EnemyTexture",
		game.AssetLoader.LoadImage("assets/img/red.png"),
		globalPosition,
	)
	enemy.AddChild(texture)

	// Add circle collider
	collider := engine.NewCircleCollider("EnemyCollider", &globalPosition.Position, 32)
	enemy.AddChild(collider)

	return enemy
}
