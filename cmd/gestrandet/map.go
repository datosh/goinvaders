package main

import (
	"engine/tml"
	"net/http"

	"github.com/hajimehoshi/ebiten"
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

func (m *Map) Width() int {
	return m.mapLoader.Width * m.mapLoader.TileWidth
}

func (m *Map) Height() int {
	return m.mapLoader.Height * m.mapLoader.TileHeight
}
