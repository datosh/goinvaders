package spaceinvaders

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type Enemy struct {
	*Sprite
	moveTimer    time.Time
	moveEach     time.Duration
	moveDistance float64
}

func NewEnemy(x, y float64, variant int) *Enemy {
	enemy := &Enemy{
		Sprite:       NewSprite(fmt.Sprintf("/sprite%d.png", variant)),
		moveTimer:    time.Now(),
		moveEach:     time.Second / 2,
		moveDistance: 20.0,
	}
	enemy.x = x
	enemy.y = y
	return enemy
}

func (e *Enemy) Update(screen *ebiten.Image) error {
	if e.moveTimer.Add(e.moveEach).Before(time.Now()) {
		e.MoveRelative(e.moveDistance, 0)
		e.moveTimer = time.Now()
	}
	return nil
}
