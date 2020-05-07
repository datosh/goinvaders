package spaceinvaders

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	player       *Player
	enemies      []*Enemy
	projectiles  []*Projectile
	fireCooldown bool
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.player.Update(screen)

	for _, enemy := range g.enemies {
		enemy.Update(screen)
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if !g.fireCooldown {
			g.projectiles = append(
				g.projectiles,
				NewProjectile(g.player.x+24, g.player.y-5),
			)

			t := time.NewTimer(time.Second / 3)
			g.fireCooldown = true
			go func() {
				<-t.C
				g.fireCooldown = false
			}()
		}
	}

	for _, projectile := range g.projectiles {
		projectile.Update(screen)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{83, 104, 138, 255})

	g.player.Draw(screen)

	for _, enemy := range g.enemies {
		enemy.Draw(screen)
	}

	for _, projectile := range g.projectiles {
		projectile.Draw(screen)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func NewGame() *Game {
	game := &Game{
		player:       NewPlayer(),
		fireCooldown: false,
	}

	game.enemies = append(game.enemies, NewEnemy(20, 20, 1))
	game.enemies = append(game.enemies, NewEnemy(120, 20, 2))
	game.enemies = append(game.enemies, NewEnemy(220, 20, 1))

	return game
}
