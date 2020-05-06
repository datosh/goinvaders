package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"net/http"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/rakyll/statik/fs"

	_ "spaceinvaders/statik"
)

var spritesFS http.FileSystem

func init() {
	var err error
	spritesFS, err = fs.New()
	if err != nil {
		log.Fatalln(err)
	}
}

func LoadImage(path string) *ebiten.Image {
	file, err := spritesFS.Open(path)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer func() {
		_ = file.Close()
	}()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Println(err)
		return nil
	}

	img2, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Println(err)
		return nil
	}
	return img2
}

type Player struct {
	img   *ebiten.Image
	x, y  float64
	scale float64
	speed float64
}

func NewPlayer() *Player {
	player := &Player{
		x:     300,
		y:     440,
		scale: 1,
		speed: 2,
	}
	player.img = LoadImage("/canon.png")

	return player
}

func (p *Player) MoveRelative(x, y float64) {
	p.x += x
	p.y += y
}

func (p *Player) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(p.scale, p.scale)
	options.GeoM.Translate(p.x, p.y)
	screen.DrawImage(p.img, options)
}

type Enemy struct {
	img   *ebiten.Image
	x, y  float64
	scale float64
	speed float64
}

func NewEnemy(x, y float64, variant int) *Enemy {
	enemy := &Enemy{
		x:     x,
		y:     y,
		scale: 1,
		speed: 2,
	}
	enemy.img = LoadImage(fmt.Sprintf("/sprite%d.png", variant))
	return enemy
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(e.scale, e.scale)
	options.GeoM.Translate(e.x, e.y)
	screen.DrawImage(e.img, options)
}

func (e *Enemy) MoveRelative(x, y float64) {
	e.x += x
	e.y += y
}

type Projectile struct {
	img   *ebiten.Image
	x, y  float64
	scale float64
	speed float64
}

func NewProjectile(x, y float64) *Projectile {
	projectile := &Projectile{
		x:     x,
		y:     y,
		scale: 1,
		speed: 2,
	}
	projectile.img = LoadImage("/projectile.png")

	return projectile
}

func (p *Projectile) Update(screen *ebiten.Image) error {
	p.y -= 1 * p.speed

	return nil
}

func (p *Projectile) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(p.scale, p.scale)
	options.GeoM.Translate(p.x, p.y)
	screen.DrawImage(p.img, options)
}

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

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
