package engine

import (
	"image"
	"io/ioutil"
	"log"
	"net/http"

	"engine/vec2"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"golang.org/x/image/font"
)

const (
	sampleRate = 48000
)

// AssetLoader can be used with any FileSystem, and provides helper functions
// to load images (png), sound (mp3), and font (ttf) files.
type AssetLoader struct {
	fs           http.FileSystem
	audioContext *audio.Context
}

// NewAssetLoader creates a new asset loader using the specified file system.
func NewAssetLoader(fs http.FileSystem) *AssetLoader {
	assetLoader := &AssetLoader{
		fs:           fs,
		audioContext: nil,
	}

	var err error
	assetLoader.audioContext, err = audio.NewContext(sampleRate)
	if err != nil {
		log.Panicf("Error initializing audio context: %v", err)
	}

	return assetLoader
}

// LoadImage from embedded filesystem.
func (loader *AssetLoader) LoadImage(path string) *ebiten.Image {
	file, err := loader.fs.Open(path)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer func() {
		_ = file.Close()
	}()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Println(err)
		return nil
	}

	img2, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Println(err)
		return nil
	}
	return img2
}

// LoadSubImage uses `LoadImage` to load an image,
// but only returns a sub image specified by bounds.
func (loader *AssetLoader) LoadSubImage(path string, bounds image.Rectangle) *ebiten.Image {
	return loader.LoadImage(path).SubImage(bounds).(*ebiten.Image)
}

// CoordinatesToBounds creates the required bounds Rectangle for `LoadSubImage`
// based on the size of the tiles, which make up a spritesheet, as well as the
// location on the sprite sheet.
func CoordinatesToBounds(tileSize vec2.I, coordinates vec2.I) image.Rectangle {
	return image.Rectangle{
		image.Point{coordinates.X * tileSize.X, coordinates.Y * tileSize.Y},
		image.Point{(coordinates.X + 1) * tileSize.X, (coordinates.Y + 1) * tileSize.Y},
	}
}

// LoadAudioPlayer created a new audio player for the specified audio file,
// which is loaded from embedded (statik) assets.
func (loader *AssetLoader) LoadAudioPlayer(path string) *audio.Player {
	f, err := loader.fs.Open(path)
	if err != nil {
		log.Panicf("Error opening audio file: %v", err)
		return nil
	}

	m, err := mp3.Decode(loader.audioContext, f)
	if err != nil {
		log.Panicf("Error decoding: %v", err)
	}

	player, err := audio.NewPlayer(loader.audioContext, m)
	if err != nil {
		log.Panicf("Error creating audio player: %v", err)
	}
	return player
}

// LoadFont loads font location at path with size in pixel.
func (loader *AssetLoader) LoadFont(path string, size float64) font.Face {
	ttfFile, err := loader.fs.Open(path)
	if err != nil {
		log.Panicf("Error opening ttf file: %v", err)
	}
	defer func() {
		err = ttfFile.Close()
		if err != nil {
			log.Panicf("Error closing ttf file: %v", err)
		}
	}()

	fontBytes, err := ioutil.ReadAll(ttfFile)
	if err != nil {
		log.Panicf("Error reading from ttfFile: %v", err)
	}

	tt, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Panicf("Error parsing font bytes: %v", err)
	}

	return truetype.NewFace(tt, &truetype.Options{
		Size:    size,
		Hinting: font.HintingFull,
	})
}
