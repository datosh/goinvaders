package main

import (
	"engine"
	"engine/vec2"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
)

type Player struct {
	*engine.Entity
	speed float64

	projectiles             []*Projectile
	fireOnCooldown          bool
	fireCooldown            time.Duration
	fireSound               *audio.Player
	fireProjectileDirection *vec2.T
}

func NewPlayer() *Player {
	player := &Player{
		Entity:                  engine.NewEntity(),
		speed:                   4,
		fireOnCooldown:          false,
		fireCooldown:            time.Second / 3,
		fireSound:               engine.LoadAudioPlayer("/audio/pew.mp3"),
		fireProjectileDirection: vec2.UY().Mul(5).Invert(),
	}
	player.fireSound.SetVolume(0.2)
	player.Image = engine.LoadSubImage(
		"/img/spritemap.png",
		engine.CoordinatesToBounds(vec2.I{64, 48}, vec2.I{2, 3}),
	)
	player.Position = &vec2.T{255, 420}
	player.ImageScale = 1.2
	player.HitboxSize = vec2.New(64, 48)

	return player
}

func (p *Player) Update(screen *ebiten.Image) error {
	p.Entity.Update(screen)

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.Position.Add(vec2.UX().Mul(p.speed).Invert())
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.Position.Add(vec2.UX().Mul(p.speed))
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if !p.fireOnCooldown {
			p.fire()
		}
	}

	for _, projectile := range p.projectiles {
		projectile.Update(screen)
	}

	p.projectiles = engine.Filter(p.projectiles, isAlive).([]*Projectile)

	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.Entity.Draw(screen)

	for _, projectile := range p.projectiles {
		projectile.Draw(screen)
	}
}

func (p *Player) CollideWith(projectile *Projectile) bool {
	return p.Hitbox().Intersects(projectile.Hitbox())
}

func (p *Player) fire() {
	p.fireSound.Rewind()
	p.fireSound.Play()
	p.addProjectile(NewProjectile(
		p.Position.Added(vec2.New(29, -7)),
		p.fireProjectileDirection,
	))
	p.startFireCooldown()
}

func (p *Player) addProjectile(projectile *Projectile) {
	p.projectiles = append(p.projectiles, projectile)
}

func (p *Player) startFireCooldown() {
	p.fireOnCooldown = true
	t := time.NewTimer(p.fireCooldown)
	go func() {
		<-t.C
		p.fireOnCooldown = false
	}()
}
