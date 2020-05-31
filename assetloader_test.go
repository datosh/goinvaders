package engine

import (
	_ "image/png"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	TestFilesDir = http.Dir("testfiles")
)

func TestNewAssetLoader(t *testing.T) {
	a := NewAssetLoader(TestFilesDir)

	t.Run("Load Blue Image", func(t *testing.T) {
		blue := a.LoadImage("blue.png")

		assert.Equal(t, 64, blue.Bounds().Size().X)
		assert.Equal(t, 64, blue.Bounds().Size().Y)
	})

	t.Run("Load MP3", func(t *testing.T) {
		mp3 := a.LoadAudioPlayer("pew.mp3")

		assert.Equal(t, int64(0), mp3.Current().Microseconds())
	})

	// This used to panic, since audio.Context can only be created once by ebiten
	t.Run("Create second asset loader", func(t *testing.T) {
		a2 := NewAssetLoader(TestFilesDir)
		blue := a2.LoadImage("blue.png")

		assert.Equal(t, 64, blue.Bounds().Size().X)
		assert.Equal(t, 64, blue.Bounds().Size().Y)
	})
}
