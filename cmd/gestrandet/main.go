package main

import (
	"embed"
	"engine"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed assets/**/*.png
	//go:embed assets/ttf/*.ttf
	//go:embed assets/config/levels/*.json
	assetFS     embed.FS
	assetLoader *engine.AssetLoader
)

func init() {
	assetLoader = engine.NewAssetLoader(assetFS)
}

func main() {
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Gestrandet")
	if err := ebiten.RunGame(NewGestrandet()); err != nil {
		log.Fatal(err)
	}
}
