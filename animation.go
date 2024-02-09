package engine

import (
	"image"
	"log"
	"time"

	"engine/vec2"

	"github.com/hajimehoshi/ebiten/v2"
)

// Animation provides logic to generate an animation from a single sprite sheet.
// All frames of the animation need to be on this sprite sheet, also all tiles
// on the sprite sheet (one tile on sprite sheet will make up one frame in
// animation) need to be of the same size.
// Animation allows to set the transition times between all frames separately.
type Animation struct {
	spritesheet     *ebiten.Image
	tileSize        vec2.I          // size of each tile on spritesheet
	tiles           []vec2.I        // list of coordinates for tiles on spritesheet
	delays          []time.Duration // delay befor going to next frame
	lastFrameChange time.Time
	currentFrame    int  // index into `tiles` & `delays` of current frame
	paused          bool // pauses animation
}

// NewAnimation builds a new animation from the provided parameters
func NewAnimation(spritesheet *ebiten.Image, tileSize vec2.I, tiles []vec2.I, delays []time.Duration) *Animation {
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
func (anim *Animation) Update() {
	if anim.paused {
		return
	}
	if time.Now().After(anim.nextFrameChange()) {
		anim.lastFrameChange = time.Now()
		anim.currentFrame = anim.nextFrame()
	}
}

// CurrentImage returns image of animation that is ought to be displayed.
func (anim *Animation) CurrentImage() *ebiten.Image {
	minX := anim.currentTile().X * anim.tileSize.X
	minY := anim.currentTile().Y * anim.tileSize.Y
	return anim.spritesheet.SubImage(image.Rect(
		minX, minY,
		minX+anim.tileSize.X, minY+anim.tileSize.Y),
	).(*ebiten.Image)
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

// Reset animation back to first frame.
func (anim *Animation) Reset() {
	anim.currentFrame = 0
	anim.lastFrameChange = time.Now()
}

func (anim *Animation) currentTile() vec2.I {
	return anim.tiles[anim.currentFrame]
}

func (anim *Animation) currentDelay() time.Duration {
	return anim.delays[anim.currentFrame]
}

func (anim *Animation) nextFrameChange() time.Time {
	return anim.lastFrameChange.Add(anim.currentDelay())
}

func (anim *Animation) nextFrame() int {
	return (anim.currentFrame + 1) % len(anim.tiles)
}

// UniformDuration is a helper function to create a slice with `n` times the
// same duration. Can be used to create a NewAnimation without repeating
// same duration multiple times.
func UniformDuration(d time.Duration, n int) []time.Duration {
	duration := make([]time.Duration, n)
	for i := range duration {
		duration[i] = d
	}
	return duration
}
