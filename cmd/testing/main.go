package main

import (
	"engine"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"

	_ "engine/statik"
	"engine/vec2"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Testing stuff!")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	spaceship *Spaceship
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.spaceship.Update(screen)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.spaceship.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

type Spaceship struct {
	*engine.Entity
}

func NewSpaceship() *Spaceship {
	spaceship := &Spaceship{
		Entity: engine.NewEntity(),
	}
	spaceship.Image = engine.LoadSubImage(
		"/img/spritemap.png",
		engine.CoordinatesToBounds(
			vec2.I{64, 48},
			vec2.I{2, 3},
		),
	)
	spaceship.Position = &vec2.T{200, 200}
	spaceship.HitboxSize = &vec2.T{51, 51}
	spaceship.HitboxOffset = &vec2.T{6, -3}
	return spaceship
}

func (s *Spaceship) Update(screen *ebiten.Image) error {
	s.Entity.Update(screen)

	// Change Position
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		s.Position.Add(vec2.UX().Invert())
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		s.Position.Add(vec2.UX())
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		s.Position.Add(vec2.UY().Invert())
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		s.Position.Add(vec2.UY())
	}

	// Change HitBox Size
	if ebiten.IsKeyPressed(ebiten.KeyJ) {
		s.HitboxSize.Add(vec2.UX().Invert())
	}
	if ebiten.IsKeyPressed(ebiten.KeyL) {
		s.HitboxSize.Add(vec2.UX())
	}
	if ebiten.IsKeyPressed(ebiten.KeyI) {
		s.HitboxSize.Add(vec2.UY().Invert())
	}
	if ebiten.IsKeyPressed(ebiten.KeyK) {
		s.HitboxSize.Add(vec2.UY())
	}

	// Change HitBox Offset
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		s.HitboxOffset.Add(vec2.UX().Invert())
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		s.HitboxOffset.Add(vec2.UX())
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		s.HitboxOffset.Add(vec2.UY().Invert())
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		s.HitboxOffset.Add(vec2.UY())
	}

	// Change Image Scale
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		s.ImageScale -= .1
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		s.ImageScale += .1
	}

	// Change Image Offset
	if ebiten.IsKeyPressed(ebiten.KeyF) {
		s.ImageOffset.Add(vec2.UX().Invert())
	}
	if ebiten.IsKeyPressed(ebiten.KeyH) {
		s.ImageOffset.Add(vec2.UX())
	}
	if ebiten.IsKeyPressed(ebiten.KeyT) {
		s.ImageOffset.Add(vec2.UY().Invert())
	}
	if ebiten.IsKeyPressed(ebiten.KeyG) {
		s.ImageOffset.Add(vec2.UY())
	}

	return nil
}

func NewGame() *Game {
	game := &Game{
		spaceship: NewSpaceship(),
	}
	return game
}
