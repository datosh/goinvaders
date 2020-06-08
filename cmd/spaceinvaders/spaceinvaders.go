package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

type Spaceinvaders struct {
	player          *Player
	enemyController *EnemyController
	stars           []*Star
	score           *Score
	levels          *Levels
	currentLevel    int
	gameOver        bool
}

func (si *Spaceinvaders) Update(screen *ebiten.Image) error {
	if si.gameOver {
		return nil
	}

	si.player.Update(screen)
	si.enemyController.Update(screen)
	si.score.Update(screen)

	for _, projectile := range si.player.projectiles {
		si.enemyController.CollideWith(projectile)
	}
	for _, projectile := range si.enemyController.projectiles {
		if si.player.CollideWith(projectile) {
			si.gameOver = true
		}
	}

	for _, star := range si.stars {
		star.Update(screen)
	}

	return nil
}

func (si *Spaceinvaders) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{21, 12, 37, 255})

	if si.gameOver {
		text.Draw(
			screen, "GAME OVER",
			assetLoader.LoadFont("/ttf/Orbitron.ttf", 32),
			180, 220,
			color.RGBA{255, 0, 0, 255},
		)
		return
	}

	for _, star := range si.stars {
		star.Draw(screen)
	}

	si.player.Draw(screen)
	si.enemyController.Draw(screen)
	si.score.Draw(screen)

	if len(si.enemyController.Enemies) == 0 {
		if si.currentLevel != si.levels.LastLevel() {
			si.currentLevel++
			si.levels.Load(si, si.currentLevel)
			si.score.SetLevel(si.currentLevel)
		}
	}
}

func (si *Spaceinvaders) Layout(int, int) (int, int) {
	return 640, 480
}

func NewSpaceinvaders() *Spaceinvaders {
	spaceinvaders := &Spaceinvaders{
		player:       NewPlayer(),
		score:        NewScore(),
		levels:       &Levels{},
		currentLevel: 1,
		gameOver:     false,
	}
	spaceinvaders.enemyController = NewEnemyController(spaceinvaders.score)
	spaceinvaders.levels.Load(spaceinvaders, spaceinvaders.currentLevel)

	for i := 0; i < 15; i++ {
		spaceinvaders.stars = append(
			spaceinvaders.stars,
			NewStar(NewStarAnimation(), spaceinvaders),
		)
	}

	return spaceinvaders
}
