package main

import (
	"time"

	"engine"
	"engine/vec2"

	"github.com/hajimehoshi/ebiten/v2"
)

type orientation int

const (
	left orientation = iota
	right
	up
	down
)

type Player struct {
	*engine.Entity
	speed          float64
	walkAnimations [4]*engine.Animation
	orientation    orientation
}

func NewPlayer() *Player {
	player := &Player{
		Entity: engine.NewEntity(),
		speed:  2.3,
	}

	spritesheet := assetLoader.LoadImage("assets/img/Character Spritemap.png")
	player.walkAnimations[down] = engine.NewAnimation(
		spritesheet,
		vec2.I{X: 40, Y: 56},
		[]vec2.I{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0}},
		engine.UniformDuration(time.Millisecond*200, 4),
	)
	player.walkAnimations[right] = engine.NewAnimation(
		spritesheet,
		vec2.I{X: 40, Y: 56},
		[]vec2.I{{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 3, Y: 1}},
		engine.UniformDuration(time.Millisecond*200, 4),
	)
	player.walkAnimations[left] = engine.NewAnimation(
		spritesheet,
		vec2.I{X: 40, Y: 56},
		[]vec2.I{{X: 0, Y: 2}, {X: 1, Y: 2}, {X: 2, Y: 2}, {X: 3, Y: 2}},
		engine.UniformDuration(time.Millisecond*200, 4),
	)
	player.walkAnimations[up] = engine.NewAnimation(
		spritesheet,
		vec2.I{X: 40, Y: 56},
		[]vec2.I{{X: 0, Y: 3}, {X: 1, Y: 3}, {X: 2, Y: 3}, {X: 3, Y: 3}},
		engine.UniformDuration(time.Millisecond*200, 4),
	)

	player.orientation = down
	player.updateAnimation()

	player.Position = &vec2.T{X: 955, Y: 720}
	player.ImageScale = 1.5
	player.HitboxSize = vec2.New(64, 48)

	return player
}

func (p *Player) Update() error {
	p.Entity.Update()
	p.updateAnimation()

	direction := vec2.New(0, 0)
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		direction.Add(vec2.UY().Invert())
		p.orientation = up
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		direction.Add(vec2.UY())
		p.orientation = down
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		direction.Add(vec2.UX().Invert())
		p.orientation = left
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		direction.Add(vec2.UX())
		p.orientation = right
	}
	direction.Normalize()
	direction.Mul(p.speed)
	p.Position.Add(direction)

	if direction.Null() {
		p.walkAnimations[p.orientation].Pause()
		p.walkAnimations[p.orientation].Reset()
	} else {
		p.walkAnimations[p.orientation].Resume()
	}

	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.Entity.Draw(screen)
}

func (p *Player) updateAnimation() {
	p.walkAnimations[p.orientation].Update()
	p.Image = p.walkAnimations[p.orientation].CurrentImage()
}
