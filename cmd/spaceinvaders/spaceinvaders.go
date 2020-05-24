package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Spaceinvaders struct {
	player          *Player
	enemyController *EnemyController
	stars           []*Star
}

func (si *Spaceinvaders) Update(screen *ebiten.Image) error {
	si.player.Update(screen)
	si.enemyController.Update(screen)

	for _, projectile := range si.player.projectiles {
		si.enemyController.CollideWith(projectile)
	}

	for _, star := range si.stars {
		star.Update(screen)
	}

	return nil
}

func (si *Spaceinvaders) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{21, 12, 37, 255})

	for _, star := range si.stars {
		star.Draw(screen)
	}

	si.player.Draw(screen)
	si.enemyController.Draw(screen)

	if len(si.enemyController.Enemies) == 0 {
		ebitenutil.DebugPrintAt(
			screen, "WINNER WINNER CHICKEN DINNER",
			640/2-100, 480/2,
		)
	}

	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("FPS %f, TPS %f", ebiten.CurrentFPS(), ebiten.CurrentTPS()),
	)
}

func (si *Spaceinvaders) Layout(int, int) (int, int) {
	return 640, 480
}

func NewSpaceinvaders() *Spaceinvaders {
	spaceinvaders := &Spaceinvaders{
		player:          NewPlayer(),
		enemyController: NewEnemyController(),
	}

	for i := 0; i < 15; i++ {
		spaceinvaders.stars = append(
			spaceinvaders.stars,
			NewStar(NewStarAnimation()),
		)
	}

	return spaceinvaders
}
