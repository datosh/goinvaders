package spaceinvaders

import (
	"reflect"

	"github.com/hajimehoshi/ebiten"
)

type Killable interface {
	Dead() bool
}

type Sprite struct {
	img   *ebiten.Image
	x, y  float64
	scale float64
	speed float64

	alive bool
}

func NewSprite(path string) *Sprite {
	sprite := &Sprite{
		img:   LoadImage(path),
		x:     0,
		y:     0,
		scale: 1,
		speed: 1,
		alive: true,
	}
	return sprite
}

func (s *Sprite) MoveRelative(x, y float64) {
	s.x += x
	s.y += y
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(s.scale, s.scale)
	options.GeoM.Translate(s.x, s.y)
	screen.DrawImage(s.img, options)
}

func (s *Sprite) Bounds() Rect {
	max := s.img.Bounds().Max
	return Rect{s.x, s.y, float64(max.X), float64(max.Y)}
}

func (s *Sprite) Dead() bool {
	return !s.alive
}

func isAlive(elem interface{}) bool {
	sprite := elem.(Killable)
	return !sprite.Dead()
}

func Filter(arr interface{}, cond func(interface{}) bool) interface{} {
	contentType := reflect.TypeOf(arr)
	contentValue := reflect.ValueOf(arr)

	newContent := reflect.MakeSlice(contentType, 0, 0)
	for i := 0; i < contentValue.Len(); i++ {
		if content := contentValue.Index(i); cond(content.Interface()) {
			newContent = reflect.Append(newContent, content)
		}
	}
	return newContent.Interface()
}
