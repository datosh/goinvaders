package main

import (
	"engine/vec2"
	"log"

	"github.com/hajimehoshi/ebiten"
)

var (
	MapSize     = vec2.NewI(23, 14)
	MapTileSize = vec2.NewI(64, 64)
)

type Map struct {
	image *ebiten.Image
}

func NewMap() *Map {
	m := &Map{}

	var err error
	m.image, err = ebiten.NewImage(
		MapSize.X*MapTileSize.X,
		MapSize.Y*MapTileSize.Y,
		ebiten.FilterDefault,
	)
	if err != nil {
		log.Fatalf("Error creating map image: %v", err)
	}

	grass := assetLoader.LoadImage("/img/grass.png")

	for x := 0; x < MapSize.X; x++ {
		for y := 0; y < MapSize.Y; y++ {
			options := &ebiten.DrawImageOptions{}
			options.GeoM.Translate(
				vec2.NewI(x, y).AsT().Mul(64.0).Coords(),
			)
			m.image.DrawImage(grass, options)
		}
	}

	return m
}

func (m *Map) Draw(screen *ebiten.Image) {
	screen.DrawImage(m.image, nil)
}
