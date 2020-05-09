package spaceinvaders

import (
	"image"
	"spaceinvaders/vec2"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type Animation struct {
	spritesheet *ebiten.Image
	tileSize    vec2.Vec2I
	tiles       []vec2.Vec2I
	delays      []time.Duration

	lastAnimChange time.Time
	currentTile    int
}

func NewAnimation(path string, tileSize vec2.Vec2I, tiles []vec2.Vec2I, delays []time.Duration) *Animation {
	anim := &Animation{
		spritesheet:    LoadImage(path),
		tileSize:       tileSize,
		tiles:          tiles,
		delays:         delays,
		lastAnimChange: time.Now(),
		currentTile:    0,
	}
	return anim
}

func (a *Animation) Update(screen *ebiten.Image) {
	if time.Now().After(a.lastAnimChange.Add(a.delays[a.currentTile])) {
		a.lastAnimChange = time.Now()
		a.currentTile = (a.currentTile + 1) % len(a.tiles)
	}
}

func (a *Animation) CurrentImage() *ebiten.Image {
	sub := a.spritesheet.SubImage(image.Rect(
		a.tiles[a.currentTile].X*a.tileSize.X,
		a.tiles[a.currentTile].Y*a.tileSize.Y,
		(a.tiles[a.currentTile].X+1)*a.tileSize.X,
		(a.tiles[a.currentTile].Y+1)*a.tileSize.Y)).(*ebiten.Image)
	return sub
}
