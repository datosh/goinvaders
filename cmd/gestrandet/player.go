package main

import (
	"engine"
	"engine/vec2"
	"time"

	"github.com/hajimehoshi/ebiten"
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

	spritesheet := assetLoader.LoadImage("/img/Character Spritemap.png")
	player.walkAnimations[down] = engine.NewAnimation(
		spritesheet,
		vec2.I{40, 56},
		[]vec2.I{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
		engine.UniformDuration(time.Millisecond*200, 4),
	)
	player.walkAnimations[right] = engine.NewAnimation(
		spritesheet,
		vec2.I{40, 56},
		[]vec2.I{{0, 1}, {1, 1}, {2, 1}, {3, 1}},
		engine.UniformDuration(time.Millisecond*200, 4),
	)
	player.walkAnimations[left] = engine.NewAnimation(
		spritesheet,
		vec2.I{40, 56},
		[]vec2.I{{0, 2}, {1, 2}, {2, 2}, {3, 2}},
		engine.UniformDuration(time.Millisecond*200, 4),
	)
	player.walkAnimations[up] = engine.NewAnimation(
		spritesheet,
		vec2.I{40, 56},
		[]vec2.I{{0, 3}, {1, 3}, {2, 3}, {3, 3}},
		engine.UniformDuration(time.Millisecond*200, 4),
	)

	player.orientation = down
	player.updateAnimation()

	player.Position = &vec2.T{255, 420}
	player.ImageScale = 1.5
	player.HitboxSize = vec2.New(64, 48)

	return player
}

func (p *Player) Update(screen *ebiten.Image) error {
	p.Entity.Update(screen)
	p.updateAnimation()

	direction := vec2.New(0, 0)
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		direction.Add(vec2.UX().Invert())
		p.orientation = left
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		direction.Add(vec2.UX())
		p.orientation = right
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		direction.Add(vec2.UY().Invert())
		p.orientation = up
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		direction.Add(vec2.UY())
		p.orientation = down
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
	p.walkAnimations[p.orientation].Update(nil)
	p.Image = p.walkAnimations[p.orientation].CurrentImage()
}
