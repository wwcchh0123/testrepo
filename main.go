package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

const (
	screenWidth   = 640
	screenHeight  = 480
	playerSpeed   = 5
	bulletSpeed   = 8
	enemySpeed    = 2
	shootCooldown = 15 // Cooldown in frames (1/4 second at 60fps)
)

var (
	playerImage *ebiten.Image
	bulletImage *ebiten.Image
	enemyImage  *ebiten.Image
)

func init() {
	// Decode images
	img, _, err := image.Decode(bytes.NewReader(Enemy_png))
	if err != nil {
		log.Fatal(err)
	}
	playerImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(Bullet_png))
	if err != nil {
		log.Fatal(err)
	}
	bulletImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(Player_png))
	if err != nil {
		log.Fatal(err)
	}
	enemyImage = ebiten.NewImageFromImage(img)
}

type Player struct {
	x, y  float64
	image *ebiten.Image
}

func (p *Player) GetRect() image.Rectangle {
	bounds := p.image.Bounds()
	return image.Rect(int(p.x), int(p.y), int(p.x)+bounds.Dx(), int(p.y)+bounds.Dy())
}

type Bullet struct {
	x, y  float64
	image *ebiten.Image
}

func (b *Bullet) GetRect() image.Rectangle {
	bounds := b.image.Bounds()
	return image.Rect(int(b.x), int(b.y), int(b.x)+bounds.Dx(), int(b.y)+bounds.Dy())
}

type Enemy struct {
	x, y  float64
	image *ebiten.Image
}

func (e *Enemy) GetRect() image.Rectangle {
	bounds := e.image.Bounds()
	return image.Rect(int(e.x), int(e.y), int(e.x)+bounds.Dx(), int(e.y)+bounds.Dy())
}

type Game struct {
	player          *Player
	bullets         []*Bullet
	enemies         []*Enemy
	enemySpawnTimer int
	shootTimer      int
	score           int
	gameOver        bool
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
	if g.gameOver {
		if ebiten.IsKeyPressed(ebiten.KeyR) {
			// Reset the game by creating a new one
			*g = *NewGame()
		}
		return nil
	}

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

	// Keep player on screen
	if g.player.x < 0 {
		g.player.x = 0
	}
	if g.player.x > screenWidth-float64(g.player.image.Bounds().Dx()) {
		g.player.x = screenWidth - float64(g.player.image.Bounds().Dx())
	}

	// Decrement shoot timer
	if g.shootTimer > 0 {
		g.shootTimer--
	}

	// Player shooting
	if ebiten.IsKeyPressed(ebiten.KeySpace) && g.shootTimer == 0 {
		g.shootTimer = shootCooldown
		bullet := &Bullet{
			x:     g.player.x + float64(g.player.image.Bounds().Dx())/2 - float64(bulletImage.Bounds().Dx())/2,
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
		enemyRect := enemy.GetRect()

		// Check collision with player
		if enemyRect.Overlaps(g.player.GetRect()) {
			g.gameOver = true
			return nil // Stop updates for this frame
		}

		// Check collision with bullets
		for j := len(g.bullets) - 1; j >= 0; j-- {
			bullet := g.bullets[j]
			if bullet.GetRect().Overlaps(enemyRect) {
				g.enemies = append(g.enemies[:i], g.enemies[i+1:]...)
				g.bullets = append(g.bullets[:j], g.bullets[j+1:]...)
				g.score++
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

	// Draw score
	scoreStr := fmt.Sprintf("Score: %d", g.score)
	text.Draw(screen, scoreStr, basicfont.Face7x13, 10, 20, color.White)

	// Draw Game Over message
	if g.gameOver {
		msg := "GAME OVER"
		subMsg := "Press 'R' to Restart"
		msgX := (screenWidth - len(msg)*7) / 2
		subMsgX := (screenWidth - len(subMsg)*7) / 2
		text.Draw(screen, msg, basicfont.Face7x13, msgX, screenHeight/2-20, color.White)
		text.Draw(screen, subMsg, basicfont.Face7x13, subMsgX, screenHeight/2, color.White)
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
