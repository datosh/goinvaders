package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type Gestrandet struct {
	m        *Map
	gameOver bool
	player   *Player
}

func (si *Gestrandet) Update(screen *ebiten.Image) error {
	if si.gameOver || ebiten.IsKeyPressed(ebiten.KeyQ) {
		return fmt.Errorf("Game Over")
	}
	si.player.Update(screen)

	return nil
}

func (si *Gestrandet) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{21, 12, 37, 255})
	si.m.Draw(screen)
	si.player.Draw(screen)
}

func (si *Gestrandet) Layout(int, int) (int, int) {
	return ebiten.ScreenSizeInFullscreen()
}

func NewGestrandet() *Gestrandet {
	gestrandet := &Gestrandet{
		m:        NewMap(),
		gameOver: false,
		player:   NewPlayer(),
	}

	return gestrandet
}
