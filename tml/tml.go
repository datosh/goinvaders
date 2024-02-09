// Package tml (tiled map loader) provides functionality to load json
// maps stored by Tiled map editor.
// Data structure is explained here:
// https://doc.mapeditor.org/en/stable/reference/json-map-format/#

package tml

import (
	"encoding/json"
	"io/fs"
	"log"
	"path/filepath"

	"engine"

	"github.com/hajimehoshi/ebiten/v2"
)

type TiledMap struct {
	Height     int `json:"height"`
	Width      int `json:"width"`
	TileHeight int `json:"tileheight"`
	TileWidth  int `json:"tilewidth"`

	Layers []*struct {
		Data []int  `json:"data"`
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"layers"`

	Tilesets []*struct {
		Name     string `json:"name"`
		FirstGID int    `json:"firstgid"`
		Tiles    []*struct {
			ID   int    `json:"id"`
			Path string `json:"image"`
			img  *ebiten.Image
		} `json:"tiles"`
	} `json:"tilesets"`

	offsetTable map[int]*ebiten.Image
}

func NewTiledMap(path string, fs fs.FS) *TiledMap {
	f, err := fs.Open(path)
	if err != nil {
		log.Printf("Unable to open %v", err)
		return nil
	}

	var dto TiledMap
	dto.offsetTable = make(map[int]*ebiten.Image)
	json.NewDecoder(f).Decode(&dto)

	loader := engine.NewAssetLoader(fs)
	base := filepath.Dir(path)
	for _, tileset := range dto.Tilesets {
		for _, tile := range tileset.Tiles {
			imgPath := filepath.ToSlash(filepath.Join(base, tile.Path))
			tile.img = loader.LoadImage(imgPath)
			dto.offsetTable[tileset.FirstGID+tile.ID] = tile.img
		}
	}

	return &dto
}

func (tm *TiledMap) Generate() *ebiten.Image {
	img := ebiten.NewImage(
		tm.Width*tm.TileWidth,
		tm.Height*tm.TileHeight,
	)

	for idx, tileType := range tm.Layers[0].Data {
		xPos := (idx % tm.Width) * tm.TileWidth
		yPos := (idx / tm.Width) * tm.TileHeight

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(xPos), float64(yPos))
		tile := tm.offsetTable[tileType]
		if tile == nil {
			log.Printf("Error: Cannot draw nil image!")
			return nil
		}
		img.DrawImage(tile, opts)
	}

	return img
}
