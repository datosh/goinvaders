package spaceinvaders

import (
	"image"
	"image/color"
	"reflect"
	"spaceinvaders/vec2"

	"github.com/hajimehoshi/ebiten"
)

type Killable interface {
	Dead() bool
}

type Sprite struct {
	position     vec2.Point    // position in the world
	image        *ebiten.Image // (optional) image to draw
	imageOffset  vec2.Point    // draw image at offset relative to position
	imageScale   float64       // scale image before drawing
	hitboxSize   *vec2.Point   // (optional) size of hitbox rel to position
	hitboxOffset *vec2.Point   // (optional) offset of hitbox relative to position
	alive        bool          // dead or alive?
	debug        bool          // debug mode
}

func NewSprite() *Sprite {
	sprite := &Sprite{
		position:     vec2.Point{X: 0.0, Y: 0.0},
		image:        nil,
		imageOffset:  vec2.Point{X: 0.0, Y: 0.0},
		imageScale:   1.0,
		hitboxSize:   nil,
		hitboxOffset: nil,
		alive:        true,
	}
	return sprite
}

// TODO: Image loading should return errors
func (s *Sprite) LoadImage(path string) {
	s.image = LoadImage(path)
}

func (s *Sprite) LoadSubImage(path string, bounds image.Rectangle) {
	img := LoadSubImage(path, bounds)
	s.image = img
}

func (s *Sprite) SetImage(image *ebiten.Image) {
	s.image = image
}

func (s *Sprite) ImageBounds() vec2.Rect {
	width, height := s.image.Size()
	return *vec2.NewRect(
		s.position.X,
		s.position.Y,
		float64(width)*s.imageScale,
		float64(height)*s.imageScale,
	)
}

func (s *Sprite) Update(screen *ebiten.Image) error {
	if ebiten.IsKeyPressed(ebiten.KeyD) && ebiten.IsKeyPressed(ebiten.KeyControl) {
		s.ToggleDebug()
	}
	return nil
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	if s.image != nil {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Scale(s.imageScale, s.imageScale)
		options.GeoM.Translate(s.position.X, s.position.Y)
		screen.DrawImage(s.image, options)
	}

	if s.debug {
		if s.hitboxSize != nil {
			DrawAABB(screen, *s.Hitbox(), color.RGBA{255, 0, 0, 255})
		}
	}
}

func (s *Sprite) MoveRelative(delta vec2.Point) {
	s.position.Add(delta)
}

func (s *Sprite) MoveRelativeX(deltaX float64) {
	s.position.X += deltaX
}

func (s *Sprite) MoveRelativeY(deltaY float64) {
	s.position.Y += deltaY
}

func (s *Sprite) MoveTo(newPosition vec2.Point) {
	s.position = newPosition
}

func (s *Sprite) HitboxSize() *vec2.Point {
	return s.hitboxSize
}

func (s *Sprite) SetHitboxSize(newSize *vec2.Point) {
	s.hitboxSize = newSize
}

func (s *Sprite) HitboxOffset() *vec2.Point {
	return s.hitboxOffset
}

func (s *Sprite) Hitbox() *vec2.Rect {
	return vec2.NewRect(
		s.position.X+s.hitboxOffset.X,
		s.position.Y+s.hitboxOffset.X,
		s.hitboxSize.X,
		s.hitboxSize.Y,
	)
}

func (s *Sprite) Die() {
	s.alive = false
}

func (s *Sprite) Dead() bool {
	return !s.alive
}

func (s *Sprite) ToggleDebug() {
	s.debug = !s.debug
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
