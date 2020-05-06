package spaceinvaders

import (
	"image"
	"log"
	"net/http"

	"github.com/hajimehoshi/ebiten"
	"github.com/rakyll/statik/fs"
)

var spritesFS http.FileSystem

func init() {
	var err error
	spritesFS, err = fs.New()
	if err != nil {
		log.Fatalln(err)
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