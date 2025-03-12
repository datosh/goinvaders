package main

import (
	"engine"
	"engine/vec2"
)

type Player struct {
	engine.BaseNode

	Health          int
	OnHealthChanged *engine.Signal[int]
}

func NewPlayer(game *engine.Game) *Player {
	player := &Player{
		BaseNode:        *engine.NewNode("Player"),
		Health:          100,
		OnHealthChanged: engine.NewSignal[int](),
	}

	globalPosition := engine.NewTransform("Position")
	globalPosition.Position = vec2.T{X: 100, Y: 100}
	globalPosition.Scale = vec2.T{X: 0.5, Y: 0.5}
	player.AddChild(globalPosition)

	texture := engine.NewTexture(
		"Body",
		game.AssetLoader.LoadImage("assets/img/green.png"),
		globalPosition,
	)

	player.AddChild(texture)

	// Add circle collider
	collider := engine.NewCircleCollider("PlayerCollider", &globalPosition.Position, 16)
	collider.OnCollision.Connect(func(other engine.Node) {
		if other.GetParent().Name() == "Enemy" {
			player.Health -= 10
			player.OnHealthChanged.Emit(player.Health)
		}

		if other.GetParent().Name() == "Square" {
			player.Health -= 1
			player.OnHealthChanged.Emit(player.Health)
		}
	})
	player.AddChild(collider)

	wasdController := engine.NewWASDController(5, &globalPosition.Position)
	player.AddChild(wasdController)

	return player
}

func (p *Player) Update() error {
	if err := p.BaseNode.Update(); err != nil {
		return err
	}

	return nil
}
