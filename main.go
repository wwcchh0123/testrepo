package main

import (
	"bytes"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

const (
	screenWidth  = 640
	screenHeight = 480
	playerSpeed  = 5
	bulletSpeed  = 8
	enemySpeed   = 2
)

var (
	playerImage *ebiten.Image
	bulletImage *ebiten.Image
	enemyImage  *ebiten.Image
)

func init() {
	// Decode images
	img, _, err := image.Decode(bytes.NewReader(images.Player_png))
	if err != nil {
		log.Fatal(err)
	}
	playerImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(images.Bullet_png))
	if err != nil {
		log.Fatal(err)
	}
	bulletImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(images.Enemy_png))
	if err != nil {
		log.Fatal(err)
	}
	enemyImage = ebiten.NewImageFromImage(img)
}

type Player struct {
	x, y  float64
	image *ebiten.Image
}

type Bullet struct {
	x, y  float64
	image *ebiten.Image
}

type Enemy struct {
	x, y  float64
	image *ebiten.Image
}

type Game struct {
	player        *Player
	bullets       []*Bullet
	enemies       []*Enemy
	enemySpawnTimer int
}

func NewGame() *Game {
	player := &Player{
		x:     screenWidth / 2,
		y:     screenHeight - 50,
		image: playerImage,
	}
	return &Game{
		player: player,
	}
}

func (g *Game) Update() error {
	// Player movement
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.player.x -= playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.player.x += playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.player.y -= playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.player.y += playerSpeed
	}

	// Player shooting
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		bullet := &Bullet{
			x:     g.player.x,
			y:     g.player.y,
			image: bulletImage,
		}
		g.bullets = append(g.bullets, bullet)
	}

	// Update bullets
	for i := len(g.bullets) - 1; i >= 0; i-- {
		bullet := g.bullets[i]
		bullet.y -= bulletSpeed
		if bullet.y < 0 {
			g.bullets = append(g.bullets[:i], g.bullets[i+1:]...)
		}
	}

	// Spawn enemies
	g.enemySpawnTimer++
	if g.enemySpawnTimer > 120 {
		g.enemySpawnTimer = 0
		enemy := &Enemy{
			x:     rand.Float64() * screenWidth,
			y:     0,
			image: enemyImage,
		}
		g.enemies = append(g.enemies, enemy)
	}

	// Update enemies
	for i := len(g.enemies) - 1; i >= 0; i-- {
		enemy := g.enemies[i]
		enemy.y += enemySpeed
		if enemy.y > screenHeight {
			g.enemies = append(g.enemies[:i], g.enemies[i+1:]...)
		}
	}

	// Collision detection
	for i := len(g.enemies) - 1; i >= 0; i-- {
		enemy := g.enemies[i]
		for j := len(g.bullets) - 1; j >= 0; j-- {
			bullet := g.bullets[j]
			if bullet.x > enemy.x && bullet.x < enemy.x+float64(enemy.image.Bounds().Dx()) &&
				bullet.y > enemy.y && bullet.y < enemy.y+float64(enemy.image.Bounds().Dy()) {
				g.enemies = append(g.enemies[:i], g.enemies[i+1:]...)
				g.bullets = append(g.bullets[:j], g.bullets[j+1:]...)
				break
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})

	// Draw player
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.player.x, g.player.y)
	screen.DrawImage(g.player.image, op)

	// Draw bullets
	for _, bullet := range g.bullets {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(bullet.x, bullet.y)
		screen.DrawImage(bullet.image, op)
	}

	// Draw enemies
	for _, enemy := range g.enemies {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(enemy.x, enemy.y)
		screen.DrawImage(enemy.image, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game := NewGame()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Plane Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}