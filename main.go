package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

const (
	screenWidth      = 480
	screenHeight     = 640
	platformWidth    = 100
	platformHeight   = 20
	playerWidth      = 32
	playerHeight     = 32
	gravity          = 0.4
	chargeRate       = 0.1
	maxCharge        = 15.0
	jumpPowerMultiplier = -1.2
)

var (
	playerImage     *ebiten.Image
	platformImage   *ebiten.Image
	backgroundColor = color.RGBA{200, 200, 255, 255}
)

// Special platform types
const (
	platformNormal = iota
	platformMusicBox
	platformConvenienceStore
	platformRubiksCube
	platformManhole
)

var platformColors = map[int]color.Color{
	platformNormal:           color.RGBA{100, 100, 100, 255}, // Grey
	platformMusicBox:         color.RGBA{255, 105, 180, 255}, // Pink
	platformConvenienceStore: color.RGBA{0, 255, 0, 255},   // Green
	platformRubiksCube:       color.RGBA{255, 255, 0, 255}, // Yellow
	platformManhole:          color.RGBA{139, 69, 19, 255},  // Brown
}

var platformScores = map[int]int{
	platformMusicBox:         30,
	platformConvenienceStore: 15,
	platformRubiksCube:       10,
	platformManhole:          5,
}

func init() {
	// Using Player_png for the player character
	img, _, err := image.Decode(bytes.NewReader(Player_png))
	if err != nil {
		log.Fatal(err)
	}
	playerImage = ebiten.NewImageFromImage(img)

	// Create a simple white image for platforms, we'll color it later
	platformImage = ebiten.NewImage(platformWidth, platformHeight)
	platformImage.Fill(color.White)
}

type Player struct {
	x, y    float64
	vx, vy  float64
	isJumping bool
	charge  float64
}

func (p *Player) GetRect() image.Rectangle {
	return image.Rect(int(p.x), int(p.y), int(p.x)+playerWidth, int(p.y)+playerHeight)
}

func (p *Player) Jump(charge float64) {
    angle := math.Pi / 3 // 60 degrees for a higher jump
    power := charge * jumpPowerMultiplier

    p.vx = math.Cos(angle) * power
    p.vy = math.Sin(angle) * power
	p.isJumping = true
}

type Platform struct {
	x, y float64
	kind int
}

func (p *Platform) GetRect() image.Rectangle {
	return image.Rect(int(p.x), int(p.y), int(p.x)+platformWidth, int(p.y)+platformHeight)
}

type Game struct {
	player         *Player
	platforms      []*Platform
	cameraY        float64
	score          int
	combo          int
	gameOver       bool
	isCharging     bool
}

func NewGame() *Game {
	g := &Game{}
	g.player = &Player{x: screenWidth/2 - playerWidth/2, y: screenHeight - platformHeight - playerHeight}
	g.platforms = append(g.platforms, &Platform{x: screenWidth/2 - platformWidth/2, y: screenHeight - platformHeight, kind: platformNormal})
	g.generatePlatforms()
	return g
}

func (g *Game) generatePlatforms() {
	for i := 0; i < 10; i++ {
		lastPlatform := g.platforms[len(g.platforms)-1]
		newX := rand.Float64()*(screenWidth-platformWidth)
		newY := lastPlatform.y - float64(rand.Intn(100)+80) // Increase min distance
		
		kind := platformNormal
		if rand.Float64() < 0.2 { // 20% chance for a special platform
			kind = rand.Intn(len(platformScores)) + 1
		}

		g.platforms = append(g.platforms, &Platform{x: newX, y: newY, kind: kind})
	}
}

func (g *Game) Update() error {
	if g.gameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			*g = *NewGame()
		}
		return nil
	}

	// Handle input
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if !g.player.isJumping {
			g.isCharging = true
		}
	}

	if g.isCharging {
		g.player.charge += chargeRate
		if g.player.charge > maxCharge {
			g.player.charge = maxCharge
		}
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if g.isCharging {
			g.isCharging = false
			g.player.Jump(g.player.charge)
			g.player.charge = 0
		}
	}

	// Player physics
	if g.player.isJumping {
		g.player.vy += gravity
		g.player.x += g.player.vx
		g.player.y += g.player.vy

		// Wall bouncing
		if g.player.x < 0 || g.player.x > screenWidth-playerWidth {
			g.player.vx *= -1
		}
	}

	// Check for landing
	if g.player.vy > 0 {
		for _, p := range g.platforms {
			playerRect := g.player.GetRect()
			platformRect := p.GetRect()

            // AABB collision check
			if playerRect.Max.X > platformRect.Min.X &&
               playerRect.Min.X < platformRect.Max.X &&
               playerRect.Max.Y > platformRect.Min.Y &&
               playerRect.Min.Y < platformRect.Max.Y {

                // Check if player was above the platform in the previous frame
                if (g.player.y + playerHeight) - g.player.vy <= p.y {
                    g.player.isJumping = false
                    g.player.vy = 0
                    g.player.vx = 0
                    g.player.y = p.y - playerHeight
                    
                    g.handleScoring(p)

                    // Generate new platforms if needed
                    if p == g.platforms[len(g.platforms)-5] {
                        g.generatePlatforms()
                    }
                    break
                }
			}
		}
	}

	// Game over condition
	if g.player.y > g.cameraY+screenHeight {
		g.gameOver = true
	}

	// Camera follow
	targetCameraY := g.player.y - screenHeight*2/3
	g.cameraY += (targetCameraY - g.cameraY) * 0.05

	return nil
}

func (g *Game) handleScoring(p *Platform) {
	platformCenter := p.x + platformWidth/2
	playerCenter := g.player.x + playerWidth/2
	dist := math.Abs(platformCenter - playerCenter)

	if dist < platformWidth/4 { // Landed near the center
		g.combo++
		g.score += g.combo * 2
	} else {
		g.combo = 1 // Reset combo to 1 for a normal hit
		g.score++
	}

	if score, ok := platformScores[p.kind]; ok {
		g.score += score
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)

	// Draw platforms
	for _, p := range g.platforms {
		platformOp := &ebiten.DrawImageOptions{}
		platformOp.GeoM.Translate(p.x, p.y-g.cameraY)
		platformOp.ColorM.ScaleWithColor(platformColors[p.kind])
		screen.DrawImage(platformImage, platformOp)
	}

	// Draw player
	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(g.player.x, g.player.y-g.cameraY)
	screen.DrawImage(playerImage, playerOp)

	// Draw UI
	scoreStr := fmt.Sprintf("Score: %d", g.score)
	text.Draw(screen, scoreStr, basicfont.Face7x13, 10, 20, color.Black)

	if g.isCharging {
		chargeStr := fmt.Sprintf("Power: %.1f", g.player.charge)
		text.Draw(screen, chargeStr, basicfont.Face7x13, 10, 40, color.Black)
	}

	if g.gameOver {
		msg := "GAME OVER"
		subMsg := "Press 'R' to Restart"
		msgX := (screenWidth - len(msg)*7) / 2
		subMsgX := (screenWidth - len(subMsg)*7) / 2
		text.Draw(screen, msg, basicfont.Face7x13, msgX, screenHeight/2-20, color.Black)
		text.Draw(screen, subMsg, basicfont.Face7x13, subMsgX, screenHeight/2, color.Black)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Jump Jump Game")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}