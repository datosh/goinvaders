package spaceinvaders

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct {
	player          *Player
	enemyController *EnemyController
	stars           []*Star
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.player.Update(screen)
	g.enemyController.Update(screen)

	for _, projectile := range g.player.projectiles {
		g.enemyController.CollideWith(projectile)
	}

	for _, star := range g.stars {
		star.Update(screen)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{21, 12, 37, 255})

	for _, star := range g.stars {
		star.Draw(screen)
	}

	g.player.Draw(screen)
	g.enemyController.Draw(screen)

	if len(g.enemyController.Enemies) == 0 {
		ebitenutil.DebugPrintAt(screen, "WINNER WINNER CHICKEN DINNER", 640/2-100, 480/2)
	}

	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("FPS %f, TPS %f", ebiten.CurrentFPS(), ebiten.CurrentTPS()),
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func NewGame() *Game {
	game := &Game{
		player:          NewPlayer(),
		enemyController: NewEnemyController(),
	}

	for i := 0; i < 15; i++ {
		game.stars = append(game.stars, NewStar(NewStarAnimation()))
	}

	return game
}
