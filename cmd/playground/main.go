package main

import (
	"embed"
	"image/color"
	"log"

	"engine"
	"engine/vec2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

var (
	//go:embed assets/img/*.png
	//go:embed assets/ttf/*.ttf
	assstFS     embed.FS
	assetLoader *engine.AssetLoader
)

func init() {
	assetLoader = engine.NewAssetLoader(assstFS)
}

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

func (g *Game) Update() error {
	g.spaceship.Update()
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
	nameFont font.Face
}

func NewSpaceship() *Spaceship {
	spaceship := &Spaceship{
		Entity: engine.NewEntity(),
	}
	spaceship.Image = assetLoader.LoadSubImage(
		"assets/img/spritemap.png",
		engine.CoordinatesToBounds(
			vec2.I{X: 64, Y: 48},
			vec2.I{X: 2, Y: 3},
		),
	)
	spaceship.nameFont = assetLoader.LoadFont("assets/ttf/Orbitron.ttf", 12)
	spaceship.Position = &vec2.T{X: 200, Y: 200}
	spaceship.HitboxSize = &vec2.T{X: 51, Y: 51}
	spaceship.HitboxOffset = &vec2.T{X: 6, Y: -3}
	return spaceship
}

func (s *Spaceship) Draw(screen *ebiten.Image) {
	s.Entity.Draw(screen)

	text.Draw(
		screen, "Spaceship", s.nameFont,
		s.Position.AsI().X, s.Position.AsI().Y+60,
		color.RGBA{255, 255, 0, 255},
	)
}

func (s *Spaceship) Update() error {
	s.Entity.Update()

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
