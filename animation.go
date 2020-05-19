package spaceinvaders

import (
	"image"
	"log"
	"spaceinvaders/vec2"
	"time"

	"github.com/hajimehoshi/ebiten"
)

// Animation provides logic to generate an animation from a single sprite sheet.
// All frames of the animation need to be on this sprite sheet, also all tiles
// on the sprite sheet (one tile on sprite sheet will make up one frame in
// animation) need to be of the same size.
// Animation allows to set the transition times between all frames separately.
type Animation struct {
	spritesheet     *ebiten.Image
	tileSize        vec2.PointI     // size of each tile on spritesheet
	tiles           []vec2.PointI   // list of coordinates for tiles on spritesheet
	delays          []time.Duration // delay befor going to next frame
	lastFrameChange time.Time
	currentFrame    int  // index into `tiles` & `delays` of current frame
	paused          bool // holds animation
}

// NewAnimation builds a new animation from the provided parameters
func NewAnimation(spritesheet *ebiten.Image, tileSize vec2.PointI, tiles []vec2.PointI, delays []time.Duration) *Animation {
	if len(tiles) != len(delays) {
		log.Printf("ERR: NewAnimation: tiles and delays need to be of same length.")
		return nil
	}
	if width, _ := spritesheet.Size(); width%tileSize.X != 0 {
		log.Printf("WARN: Spritesheet width isn't multiple of tilesize width.")
	}
	if _, height := spritesheet.Size(); height&tileSize.Y != 0 {
		log.Printf("WARN: Spritesheet height isn't multiple of tilesize height.")
	}

	anim := &Animation{
		spritesheet:     spritesheet,
		tileSize:        tileSize,
		tiles:           tiles,
		delays:          delays,
		lastFrameChange: time.Now(),
		currentFrame:    0,
		paused:          false,
	}
	return anim
}

// Update needs to be called regularly so that animation progresses.
func (anim *Animation) Update(_ *ebiten.Image) {
	if anim.paused {
		return
	}
	if time.Now().After(anim.nextFrameChange()) {
		anim.lastFrameChange = time.Now()
		anim.currentFrame = (anim.currentFrame + 1) % len(anim.tiles)
	}
}

// CurrentImage returns image of animation that is ought to be displayed.
func (anim *Animation) CurrentImage() *ebiten.Image {
	sub := anim.spritesheet.SubImage(image.Rect(
		anim.currentTile().X*anim.tileSize.X,
		anim.currentTile().Y*anim.tileSize.Y,
		(anim.currentTile().X+1)*anim.tileSize.X,
		(anim.currentTile().Y+1)*anim.tileSize.Y)).(*ebiten.Image)
	return sub
}

// Pause going through single frames that make up animation.
func (anim *Animation) Pause() {
	anim.paused = true
}

// Resume going through frames. Time that was spend paused is not counted, so
// transition to next frame is probably instant.
func (anim *Animation) Resume() {
	anim.paused = false
}

func (anim *Animation) currentTile() vec2.PointI {
	return anim.tiles[anim.currentFrame]
}

func (anim *Animation) currentDelay() time.Duration {
	return anim.delays[anim.currentFrame]
}

func (anim *Animation) nextFrameChange() time.Time {
	return anim.lastFrameChange.Add(anim.currentDelay())
}
