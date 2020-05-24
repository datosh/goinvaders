package spaceinvaders

import (
	"image"
	"log"
	"net/http"
	"os"

	// Import to be embedded images
	_ "spaceinvaders/statik"
	"spaceinvaders/vec2"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/rakyll/statik/fs"
)

const (
	sampleRate = 48000
)

var (
	spritesFS    http.FileSystem
	audioContext *audio.Context
)

func init() {
	initStatik()
	initAudio()
}

func logFilePath(path string, info os.FileInfo, err error) error {
	log.Println(path)
	return nil
}

func listFiles() {
	log.Println("Assets:")
	fs.Walk(spritesFS, "/", logFilePath)
}

func initStatik() {
	var err error
	spritesFS, err = fs.New()
	if err != nil {
		log.Fatalln(err)
	}
	listFiles()
}

func initAudio() {
	var err error
	audioContext, err = audio.NewContext(sampleRate)
	if err != nil {
		log.Panicf("Error initializing audio context: %v", err)
	}
}

func LoadImage(path string) *ebiten.Image {
	file, err := spritesFS.Open(path)
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

func LoadSubImage(path string, bounds image.Rectangle) *ebiten.Image {
	img := LoadImage(path)
	subImg := img.SubImage(bounds)
	finalImg := subImg.(*ebiten.Image)
	return finalImg
}

func CoordinatesToBounds(tileSize vec2.I, coordinates vec2.I) image.Rectangle {
	return image.Rectangle{
		image.Point{coordinates.X * tileSize.X, coordinates.Y * tileSize.Y},
		image.Point{(coordinates.X + 1) * tileSize.X, (coordinates.Y + 1) * tileSize.Y},
	}
}

func LoadAudioPlayer(path string) *audio.Player {
	f, err := spritesFS.Open(path)
	if err != nil {
		log.Panicf("Error opening audio file: %v", err)
		return nil
	}

	m, err := mp3.Decode(audioContext, f)
	if err != nil {
		log.Panicf("Error decoding: %v", err)
	}

	player, err := audio.NewPlayer(audioContext, m)
	if err != nil {
		log.Panicf("Error creating audio player: %v", err)
	}
	return player
}
