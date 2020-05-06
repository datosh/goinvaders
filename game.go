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
	moveTimer    time.Time
	moveEach     time.Duration
}

func (g *Game) Update(screen *ebiten.Image) error {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.player.MoveRelative(-1*g.player.speed, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.player.MoveRelative(1*g.player.speed, 0)
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if !g.fireCooldown {
			g.projectiles = append(
				g.projectiles,
				NewProjectile(g.player.x+24, g.player.y-5),
			)

			t := time.NewTimer(time.Second)
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

	if g.moveTimer.Add(g.moveEach).Before(time.Now()) {
		for _, enemy := range g.enemies {
			enemy.MoveRelative(10, 0)
		}
		g.moveTimer = time.Now()
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
		moveTimer:    time.Now(),
		moveEach:     time.Second,
		fireCooldown: false,
	}

	game.enemies = append(game.enemies, NewEnemy(20, 20, 1))
	game.enemies = append(game.enemies, NewEnemy(120, 20, 2))
	game.enemies = append(game.enemies, NewEnemy(220, 20, 1))

	return game
}
