package spaceinvaders

import (
	"image/color"
	"reflect"
	"spaceinvaders/vec2"

	"github.com/hajimehoshi/ebiten"
)

// Entity is basis for all entities in the game
type Entity struct {
	Position     *vec2.T       // position in the world
	Image        *ebiten.Image // (optional) image to draw
	ImageOffset  *vec2.T       // draw image at offset relative to position
	ImageScale   float64       // scale image before drawing
	HitboxSize   *vec2.T       // (optional) size of hitbox relative to position
	HitboxOffset *vec2.T       // offset of hitbox relative to position
	Alive        bool          // dead or alive?
	Debug        bool          // debug mode
}

// NewEntity created a new entity and sets sane defaults
func NewEntity() *Entity {
	entity := &Entity{
		Position:     &vec2.T{X: 0.0, Y: 0.0},
		Image:        nil,
		ImageOffset:  &vec2.T{X: 0.0, Y: 0.0},
		ImageScale:   1.0,
		HitboxSize:   nil,
		HitboxOffset: &vec2.T{X: 0.0, Y: 0.0},
		Alive:        true,
		Debug:        false,
	}
	return entity
}

// ImageRect builds destination rectangle of image in world
func (entity *Entity) ImageRect() vec2.Rect {
	imageSize := vec2.NewI(entity.Image.Size()).AsT()
	minPoint := entity.Position.Added(entity.ImageOffset)
	return vec2.Rect{
		Min: minPoint,
		Max: minPoint.Added(imageSize.Muled(entity.ImageScale)),
	}
}

func (s *Entity) Update(screen *ebiten.Image) error {
	if ebiten.IsKeyPressed(ebiten.KeyF10) {
		s.ToggleDebug()
	}
	return nil
}

func (s *Entity) Draw(screen *ebiten.Image) {
	if s.Image != nil {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Scale(s.ImageScale, s.ImageScale)
		options.GeoM.Translate(
			s.Position.Added(s.ImageOffset).Coords(),
		)
		screen.DrawImage(s.Image, options)
	}

	if s.Debug {
		if s.HitboxSize != nil {
			DrawRect(screen, *s.Hitbox(), color.RGBA{255, 0, 0, 255})
		}

		if s.Image != nil {
			DrawRect(screen, s.ImageRect(), color.RGBA{0, 255, 0, 255})
		}

		DrawRect(
			screen,
			*vec2.NewRect(s.Position.X, s.Position.Y, 2, 2),
			color.RGBA{0, 0, 255, 255},
		)
	}
}

func (s *Entity) Hitbox() *vec2.Rect {
	if s.HitboxSize == nil {
		return nil
	}
	return &vec2.Rect{
		Min: s.Position.Added(s.HitboxOffset),
		Max: s.Position.Added(s.HitboxOffset).Add(s.HitboxSize),
	}
}

func (s *Entity) Die() {
	s.Alive = false
}

func (s *Entity) Dead() bool {
	return !s.Alive
}

func (s *Entity) ToggleDebug() {
	s.Debug = !s.Debug
}

// TODO: Move this to some filter / container file
type Killable interface {
	Dead() bool
}

func isAlive(elem interface{}) bool {
	entity := elem.(Killable)
	return !entity.Dead()
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
