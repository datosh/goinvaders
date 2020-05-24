package main

import (
	_ "image/png"
	"log"
	"spaceinvaders"

	"github.com/hajimehoshi/ebiten"

	_ "spaceinvaders/statik"
	"spaceinvaders/vec2"
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

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.spaceship.MoveRelativeX(-1.0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.spaceship.MoveRelativeX(1.0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.spaceship.MoveRelativeY(-1.0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.spaceship.MoveRelativeY(1.0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyJ) {
		g.spaceship.SetHitboxSize(&vec2.T{
			g.spaceship.Hitbox().Width() - 1,
			g.spaceship.Hitbox().Height(),
		})
	}
	if ebiten.IsKeyPressed(ebiten.KeyL) {
		g.spaceship.SetHitboxSize(&vec2.T{
			g.spaceship.Hitbox().Width() + 1,
			g.spaceship.Hitbox().Height(),
		})
	}
	if ebiten.IsKeyPressed(ebiten.KeyI) {
		g.spaceship.SetHitboxSize(&vec2.T{
			g.spaceship.Hitbox().Width(),
			g.spaceship.Hitbox().Height() - 1,
		})
	}
	if ebiten.IsKeyPressed(ebiten.KeyK) {
		g.spaceship.SetHitboxSize(&vec2.T{
			g.spaceship.Hitbox().Width(),
			g.spaceship.Hitbox().Height() + 1,
		})
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.spaceship.SetHitboxOffset(&vec2.T{
			g.spaceship.HitboxOffset().X,
			g.spaceship.HitboxOffset().Y - 1,
		})
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.spaceship.SetHitboxOffset(&vec2.T{
			g.spaceship.HitboxOffset().X,
			g.spaceship.HitboxOffset().Y + 1,
		})
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.spaceship.SetHitboxOffset(&vec2.T{
			g.spaceship.HitboxOffset().X - 1,
			g.spaceship.HitboxOffset().Y,
		})
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.spaceship.SetHitboxOffset(&vec2.T{
			g.spaceship.HitboxOffset().X + 1,
			g.spaceship.HitboxOffset().Y,
		})
	}

	if ebiten.IsKeyPressed(ebiten.KeyQ) {

	}

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
	*spaceinvaders.Sprite
}

func NewSpaceship() *Spaceship {
	spaceship := &Spaceship{
		Sprite: spaceinvaders.NewSprite(),
	}
	spaceship.SetImage(spaceinvaders.LoadSubImage(
		"/img/spritemap.png",
		spaceinvaders.CoordinatesToBounds(
			vec2.I{64, 48},
			vec2.I{2, 3},
		),
	))
	spaceship.SetPosition(&vec2.T{200, 200})
	spaceship.SetHitboxSize(&vec2.T{51, 51})
	spaceship.SetHitboxOffset(&vec2.T{6, -3})
	return spaceship
}

func NewGame() *Game {
	game := &Game{
		spaceship: NewSpaceship(),
	}
	return game
}
