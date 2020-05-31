package main

import (
	"engine/tml"
	"engine/vec2"
	"net/http"

	"github.com/hajimehoshi/ebiten"
)

var (
	MapSize     = vec2.NewI(23, 14)
	MapTileSize = vec2.NewI(64, 64)
)

type Map struct {
	mapLoader *tml.TiledMap
	image     *ebiten.Image
}

func NewMap() *Map {
	m := &Map{
		mapLoader: tml.NewTiledMap("config/levels/gestrandet.json", http.Dir("assets")),
	}
	m.image = m.mapLoader.Generate()
	return m
}

func (m *Map) Draw(screen *ebiten.Image) {
	screen.DrawImage(m.image, nil)
}
