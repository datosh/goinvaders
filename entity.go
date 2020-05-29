package engine

import (
	"engine/vec2"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
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

// Update should be called in game loop
func (entity *Entity) Update(screen *ebiten.Image) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyF10) {
		entity.ToggleDebug()
	}
	return nil
}

// Draw should be called in draw loop
func (entity *Entity) Draw(screen *ebiten.Image) {
	if entity.Image != nil {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Scale(entity.ImageScale, entity.ImageScale)
		options.GeoM.Translate(
			entity.Position.Added(entity.ImageOffset).Coords(),
		)
		screen.DrawImage(entity.Image, options)
	}

	if entity.Debug {
		if entity.HitboxSize != nil {
			DrawRect(screen, *entity.Hitbox(), color.RGBA{255, 0, 0, 255})
		}

		if entity.Image != nil {
			DrawRect(screen, entity.ImageRect(), color.RGBA{0, 255, 0, 255})
		}

		DrawRect(
			screen,
			*vec2.NewRect(entity.Position.X, entity.Position.Y, 2, 2),
			color.RGBA{0, 0, 255, 255},
		)
	}
}

// Hitbox returns entities hitbox, which is based on players position, as well
// as hitbox offset and size
func (entity *Entity) Hitbox() *vec2.Rect {
	if entity.HitboxSize == nil {
		return nil
	}
	return &vec2.Rect{
		Min: entity.Position.Added(entity.HitboxOffset),
		Max: entity.Position.Added(entity.HitboxOffset).Add(entity.HitboxSize),
	}
}

// Die kills the entity
func (entity *Entity) Die() {
	entity.Alive = false
}

// Dead returns true if the entity is dead
func (entity *Entity) Dead() bool {
	return !entity.Alive
}

// ToggleDebug mode for this entity
func (entity *Entity) ToggleDebug() {
	entity.Debug = !entity.Debug
}
