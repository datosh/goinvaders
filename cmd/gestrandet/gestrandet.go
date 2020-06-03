package main

import (
	"engine"
	"engine/vec2"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
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

	camPos := si.player.Position.Subed(
		vec2.NewI(screen.Size()).AsT().Mul(0.5).Sub(
			vec2.NewI(si.player.Image.Size()).AsT().Mul(0.5),
		),
	)
	si.camera.MoveTo(camPos)

	if ebiten.IsKeyPressed(ebiten.KeyL) {
		si.camera.Zoom(1.01)
	}
	if ebiten.IsKeyPressed(ebiten.KeyK) {
		si.camera.Zoom(0.95)
	}

	if ebiten.IsKeyPressed(ebiten.KeyO) {
		si.camera.Rotate(0.2)
	}
	if ebiten.IsKeyPressed(ebiten.KeyP) {
		si.camera.Rotate(-0.2)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		si.camera.Reset()
	}

	return nil
}

func (si *Gestrandet) Draw(screen *ebiten.Image) {
	si.world.Fill(color.RGBA{21, 12, 37, 255})
	si.m.Draw(si.world)
	si.player.Draw(si.world)
	si.camera.Render(si.world, screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %v, FPS: %v", ebiten.CurrentTPS(), ebiten.CurrentFPS()))
	ebitenutil.DebugPrintAt(screen, si.camera.String(), 0, 20)
	mouseInfo := fmt.Sprintf(
		"OnScreen: %v, World: %v",
		vec2.NewI(ebiten.CursorPosition()),
		si.camera.ScreenToWorld(vec2.NewI(ebiten.CursorPosition()).AsT()),
	)
	ebitenutil.DebugPrintAt(screen, mouseInfo, 0, 40)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Origin of world shown on screen at: %v", si.camera.WorldToScreen(vec2.New(0, 0))), 0, 60)
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
	gestrandet.camera = engine.NewCamera(vec2.NewI(gestrandet.Layout(0, 0)).AsT())
	gestrandet.world = gestrandet.m.mapLoader.Generate()

	return gestrandet
}
