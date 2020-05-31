package tml

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	TestFilesDir = http.Dir("../testfiles")
)

func TestNewTiledMap(t *testing.T) {

	t.Run("Basic Checks", func(t *testing.T) {
		dto := NewTiledMap("testmap.json", TestFilesDir)

		assert.Equal(t, 10, dto.Width)
		assert.Equal(t, 12, dto.Height)

		assert.Equal(t, 120, len(dto.Layers[0].Data))
	})

	t.Run("Wrong path returns nil", func(t *testing.T) {
		dto := NewTiledMap("xxxxx.json", TestFilesDir)
		assert.Nil(t, dto)
	})

}

// func TestGenerate(t *testing.T) {
// 	dto := NewTiledMap("testmap.json", TestFilesDir)
// 	m := dto.Generate()

// 	t.Run("Check dimensions", func(t *testing.T) {
// 		assert.NotNil(t, m)

// 		assert.Equal(t, 640, m.Bounds().Size().X)
// 		assert.Equal(t, 768, m.Bounds().Size().Y)
// 	})

// 	f, _ := os.Create("testmap.png")
// 	png.Encode(f, m)
// }
