//go:generate statik -src=assets -include=*.png,*.mp3,*.ttf,*.json

package main

import (
	"engine"
	_ "image/png"

	_ "engine/cmd/gestrandet/statik"

	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/rakyll/statik/fs"
)

var (
	assetLoader *engine.AssetLoader
)

func init() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	assetLoader = engine.NewAssetLoader(statikFS)
}

func main() {
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Gestrandet")
	if err := ebiten.RunGame(NewGestrandet()); err != nil {
		log.Fatal(err)
	}
}
