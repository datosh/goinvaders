package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"

	_ "engine/statik"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Spaceinvaders")
	if err := ebiten.RunGame(NewSpaceinvaders()); err != nil {
		log.Fatal(err)
	}
}
