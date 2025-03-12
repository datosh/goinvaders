package engine

import (
	"io/fs"

	"engine/vec2"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Scene       *Scene
	ScreenSize  vec2.T
	Fullscreen  bool
	Title       string
	AssetLoader *AssetLoader
}

func NewGame(title string, assets fs.FS) *Game {
	return &Game{
		Scene:       NewScene(),
		ScreenSize:  vec2.T{X: 640, Y: 480},
		Fullscreen:  false,
		Title:       title,
		AssetLoader: NewAssetLoader(assets),
	}
}

func (g *Game) Run() error {
	ebiten.SetWindowTitle(g.Title)
	ebiten.SetWindowSize(int(g.ScreenSize.X), int(g.ScreenSize.Y))
	return ebiten.RunGame(g)
}

func (g *Game) Update() error {
	return g.Scene.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Scene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(g.ScreenSize.X), int(g.ScreenSize.Y)
}
