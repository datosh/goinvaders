package main

import (
	"embed"
	"engine"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed assets/audio/*.mp3
	//go:embed assets/img/*.png
	//go:embed assets/ttf/*.ttf
	assetFS     embed.FS
	assetLoader *engine.AssetLoader
)

func init() {
	assetLoader = engine.NewAssetLoader(assetFS)
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Spaceinvaders")
	if err := ebiten.RunGame(NewSpaceinvaders()); err != nil {
		log.Fatal(err)
	}
}
