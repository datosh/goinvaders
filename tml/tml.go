// Package tml (tiled map loader) provides functionality to load json
// maps stored by Tiled map editor.
// Data structure is explained here:
// https://doc.mapeditor.org/en/stable/reference/json-map-format/#

package tml

import (
	"encoding/json"
	"engine"
	"log"
	"net/http"
	"path/filepath"

	"github.com/hajimehoshi/ebiten"
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
}

func NewTiledMap(path string, fs http.FileSystem) *TiledMap {
	f, err := fs.Open(path)
	if err != nil {
		log.Panicf("Unable to open %v", err)
	}

	var dto TiledMap
	json.NewDecoder(f).Decode(&dto)

	loader := engine.NewAssetLoader(fs)
	base := filepath.Dir(path)
	for _, tileset := range dto.Tilesets {
		for _, tile := range tileset.Tiles {
			imgPath := filepath.ToSlash(filepath.Join(base, tile.Path))
			log.Printf("Loading: %v", imgPath)
			tile.img = loader.LoadImage(imgPath)
		}
	}

	return &dto
}

func (tm *TiledMap) Generate() *ebiten.Image {
	img, err := ebiten.NewImage(
		tm.Width*tm.TileWidth,
		tm.Height*tm.TileHeight,
		ebiten.FilterDefault,
	)
	if err != nil {
		log.Printf("Error creating new image for map: %v", err)
		return nil
	}

	return img
}
