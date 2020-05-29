//go:generate statik -src=assets -include=*.png,*.mp3,*.ttf

package main

import (
	"engine"
	_ "image/png"

	_ "engine/cmd/spaceinvaders/statik"

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
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Spaceinvaders")
	if err := ebiten.RunGame(NewSpaceinvaders()); err != nil {
		log.Fatal(err)
	}
}
