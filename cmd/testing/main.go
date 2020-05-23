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
	spaceship.LoadSubImage(
		"/img/spritemap.png",
		spaceinvaders.CoordinatesToBounds(
			vec2.PointI{64, 48},
			vec2.PointI{2, 3},
		),
	)
	spaceship.MoveTo(vec2.Point{200, 200})
	spaceship.SetHitboxSize(&vec2.Point{64, 48})
	return spaceship
}

func NewGame() *Game {
	game := &Game{
		spaceship: NewSpaceship(),
	}
	return game
}
