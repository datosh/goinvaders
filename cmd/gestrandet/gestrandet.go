package main

import (
	"engine"
	"engine/vec2"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Gestrandet struct {
	m        *Map
	gameOver bool
	player   *Player
	world    *ebiten.Image
	camera   *engine.Camera
}

func (si *Gestrandet) Update(screen *ebiten.Image) error {
	if si.gameOver || ebiten.IsKeyPressed(ebiten.KeyQ) {
		return fmt.Errorf("Game Over")
	}
	si.player.Update(screen)

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		si.camera.Move(vec2.UY())
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		si.camera.Move(vec2.UY().Invert())
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		si.camera.Move(vec2.UX().Invert())
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		si.camera.Move(vec2.UX())
	}

	return nil
}

func (si *Gestrandet) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{21, 12, 37, 255})
	si.m.Draw(si.world)
	si.player.Draw(si.world)

	si.camera.View(si.world, screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %v, FPS: %v", ebiten.CurrentTPS(), ebiten.CurrentFPS()))
}

func (si *Gestrandet) Layout(int, int) (int, int) {
	return ebiten.ScreenSizeInFullscreen()
}

func NewGestrandet() *Gestrandet {
	gestrandet := &Gestrandet{
		m:        NewMap(),
		gameOver: false,
		player:   NewPlayer(),
		camera:   engine.NewCamera(100, 100),
	}
	gestrandet.world = gestrandet.m.mapLoader.Generate()

	return gestrandet
}
