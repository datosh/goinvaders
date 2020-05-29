package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"

	_ "engine/statik"
)

func main() {
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Gestrandet")
	if err := ebiten.RunGame(NewGestrandet()); err != nil {
		log.Fatal(err)
	}
}
