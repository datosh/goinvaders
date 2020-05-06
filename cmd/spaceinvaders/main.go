package main

import (
	_ "image/png"
	"log"
	"spaceinvaders"

	"github.com/hajimehoshi/ebiten"

	_ "spaceinvaders/statik"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(spaceinvaders.NewGame()); err != nil {
		log.Fatal(err)
	}
}
