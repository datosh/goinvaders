package spaceinvaders

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct {
	player          *Player
	enemyController *EnemyController
	projectiles     []*Projectile
	stars           []*Star
	fireCooldown    bool
	pewPlayer       *audio.Player
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.player.Update(screen)

	g.enemyController.Update(screen)
	for _, projectile := range g.projectiles {
		projectile.Update(screen)
	}

	for _, projectile := range g.projectiles {
		g.enemyController.CollideWith(projectile)
	}

	g.projectiles = Filter(g.projectiles, isAlive).([]*Projectile)

	for _, star := range g.stars {
		star.Update(screen)
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if !g.fireCooldown {

			g.pewPlayer.SetVolume(0.2)
			g.pewPlayer.Rewind()
			g.pewPlayer.Play()
			g.projectiles = append(
				g.projectiles,
				NewProjectile(g.player.Position.Copy()),
			)
			t := time.NewTimer(time.Second / 3)
			g.fireCooldown = true
			go func() {
				<-t.C
				g.fireCooldown = false
			}()
		}
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

	for _, projectile := range g.projectiles {
		projectile.Draw(screen)
	}

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
		fireCooldown:    false,
		pewPlayer:       LoadAudioPlayer("/audio/pew.mp3"),
	}

	for i := 0; i < 15; i++ {
		game.stars = append(game.stars, NewStar(NewStarAnimation()))
	}
	return game
}
