package spaceinvaders

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct {
	player       *Player
	enemies      []*Enemy
	projectiles  []*Projectile
	stars        []*Star
	fireCooldown bool
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.player.Update(screen)

	for _, enemy := range g.enemies {
		enemy.Update(screen)
	}
	g.enemies = Filter(g.enemies, isAlive).([]*Enemy)

	for _, projectile := range g.projectiles {
		projectile.Update(screen)
	}
	g.projectiles = Filter(g.projectiles, isAlive).([]*Projectile)

	for _, star := range g.stars {
		star.Update(screen)
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if !g.fireCooldown {
			g.projectiles = append(
				g.projectiles,
				NewProjectile(g.player.x+30, g.player.y-10),
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

	for _, enemy := range g.enemies {
		enemy.Draw(screen)
	}

	for _, projectile := range g.projectiles {
		projectile.Draw(screen)
	}

	for _, enemy := range g.enemies {
		for _, projectile := range g.projectiles {
			if DoCollide(enemy.Bounds(), projectile.Bounds()) {
				enemy.alive = false
				projectile.alive = false
			}
		}
	}

	if len(g.enemies) == 0 {
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
		player:       NewPlayer(),
		fireCooldown: false,
	}

	game.enemies = append(game.enemies, NewEnemy(20, 20, NewEnemy1Animation()))
	game.enemies = append(game.enemies, NewEnemy(120, 20, NewEnemy2Animation()))
	game.enemies = append(game.enemies, NewEnemy(220, 20, NewEnemy1Animation()))

	for i := 0; i < 15; i++ {
		game.stars = append(game.stars, NewStar(NewStarAnimation()))
	}
	return game
}
