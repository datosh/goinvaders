package tml

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTiledMap(t *testing.T) {
	dto := NewTiledMap("testfiles/testmap.json", http.Dir(".."))

	assert.Equal(t, 12, dto.Height)
	assert.Equal(t, 1, dto.Layers[0].Data[0])
	assert.Equal(t, 0, dto.Tilesets[0].Tiles[0].ID)
	assert.Equal(t, 1, dto.Tilesets[0].FirstGID)
	assert.NotNil(t, dto.Tilesets[0].Tiles[0].img)
	assert.Equal(t, 64, dto.Tilesets[0].Tiles[0].img.Bounds().Dx())
}

func TestGenerate(t *testing.T) {
	// dto := NewTiledMap("testfiles/testmap.json", http.Dir(".."))
	// m := dto.Generate()

	// size := m.Bounds().Min
	// assert.Equal(t, 640, size.X)
}
