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
	screenWidth         = 480
	screenHeight        = 640
	platformWidth       = 100
	platformHeight      = 20
	playerWidth         = 32
	playerHeight        = 32
	gravity             = 0.4
	chargeRate          = 0.1
	maxCharge           = 15.0
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

// Power-up types
const (
	powerUpSpeedBoost = iota
	powerUpDoubleScore
	powerUpExtraJump
	powerUpSlowMotion
)

var platformColors = map[int]color.Color{
	platformNormal:           color.RGBA{100, 100, 100, 255}, // Grey
	platformMusicBox:         color.RGBA{255, 105, 180, 255}, // Pink
	platformConvenienceStore: color.RGBA{0, 255, 0, 255},     // Green
	platformRubiksCube:       color.RGBA{255, 255, 0, 255},   // Yellow
	platformManhole:          color.RGBA{139, 69, 19, 255},   // Brown
}

var platformScores = map[int]int{
	platformMusicBox:         30,
	platformConvenienceStore: 15,
	platformRubiksCube:       10,
	platformManhole:          5,
}

var powerUpColors = map[int]color.Color{
	powerUpSpeedBoost:  color.RGBA{255, 0, 0, 255},   // Red
	powerUpDoubleScore: color.RGBA{255, 215, 0, 255}, // Gold
	powerUpExtraJump:   color.RGBA{0, 255, 255, 255}, // Cyan
	powerUpSlowMotion:  color.RGBA{128, 0, 128, 255}, // Purple
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
	x, y      float64
	vx, vy    float64
	isJumping bool
	charge    float64
	direction float64
	// Power-up effects
	speedBoost    float64
	doubleScore   float64
	extraJumps    int
	slowMotion    float64
	powerUpTimers map[int]float64
}

func (p *Player) GetRect() image.Rectangle {
	return image.Rect(int(p.x), int(p.y), int(p.x)+playerWidth, int(p.y)+playerHeight)
}

func (p *Player) Jump(charge float64) {
	angle := math.Pi / 3 // 60 degrees for a higher jump
	power := charge * jumpPowerMultiplier

	p.vx = math.Cos(angle) * power * p.direction
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

type PowerUp struct {
	x, y      float64
	kind      int
	collected bool
}

func (p *PowerUp) GetRect() image.Rectangle {
	size := 20
	return image.Rect(int(p.x), int(p.y), int(p.x)+size, int(p.y)+size)
}

type Particle struct {
	x, y   float64
	vx, vy float64
	life   float64
	color  color.Color
}

type Game struct {
	player          *Player
	platforms       []*Platform
	powerUps        []*PowerUp
	cameraY         float64
	score           int
	highScore       int
	combo           int
	gameOver        bool
	isCharging      bool
	particles       []*Particle
	difficultyLevel int
}

func NewGame() *Game {
	g := &Game{}
	g.player = &Player{
		x:             screenWidth/2 - playerWidth/2,
		y:             screenHeight - platformHeight - playerHeight,
		direction:     1,
		powerUpTimers: make(map[int]float64),
	}
	g.platforms = append(g.platforms, &Platform{x: screenWidth/2 - platformWidth/2, y: screenHeight - platformHeight, kind: platformNormal})
	g.generatePlatforms()
	g.loadHighScore()
	return g
}

func (g *Game) generatePlatforms() {
	// Update difficulty based on score
	g.difficultyLevel = g.score / 100 // Increase difficulty every 100 points

	for i := 0; i < 10; i++ {
		lastPlatform := g.platforms[len(g.platforms)-1]

		// Difficulty affects platform spacing
		baseDist := 80 + rand.Intn(100)
		difficultyBonus := g.difficultyLevel * 10 // Increase distance with difficulty
		if difficultyBonus > 50 {                 // Cap the bonus
			difficultyBonus = 50
		}

		newX := rand.Float64() * (screenWidth - platformWidth)
		newY := lastPlatform.y - float64(baseDist+difficultyBonus)

		kind := platformNormal
		specialChance := 0.2 + float64(g.difficultyLevel)*0.05 // More special platforms at higher difficulty
		if specialChance > 0.4 {                               // Cap at 40%
			specialChance = 0.4
		}
		if rand.Float64() < specialChance {
			kind = rand.Intn(len(platformScores)) + 1
		}

		g.platforms = append(g.platforms, &Platform{x: newX, y: newY, kind: kind})

		// Generate power-ups randomly (fewer at higher difficulty)
		powerUpChance := 0.15 - float64(g.difficultyLevel)*0.01
		if powerUpChance < 0.05 { // Minimum 5% chance
			powerUpChance = 0.05
		}
		if rand.Float64() < powerUpChance {
			powerUpX := rand.Float64() * (screenWidth - 20)
			powerUpY := newY - 30
			powerUpKind := rand.Intn(4) // 4 types of power-ups
			g.powerUps = append(g.powerUps, &PowerUp{x: powerUpX, y: powerUpY, kind: powerUpKind})
		}
	}
}

func (g *Game) Update() error {
	if g.gameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			*g = *NewGame()
		}
		return nil
	}

	// Handle input (Enhanced for mobile/touch)
	touches := ebiten.AppendTouchIDs(nil)
	mousePressed := inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
	mouseReleased := inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft)
	touchPressed := len(touches) > 0 && inpututil.IsTouchJustPressed(touches[0])
	touchReleased := len(inpututil.AppendJustReleasedTouchIDs(nil)) > 0

	// Start charging on mouse/touch press
	if (mousePressed || touchPressed) && !g.player.isJumping {
		g.isCharging = true
	}

	// Continue charging while held
	if g.isCharging {
		g.player.charge += chargeRate
		if g.player.charge > maxCharge {
			g.player.charge = maxCharge
		}
	}

	// Release jump on mouse/touch release
	if (mouseReleased || touchReleased) && g.isCharging {
		g.isCharging = false

		// Apply speed boost power-up to jump
		jumpPower := g.player.charge
		if g.player.speedBoost > 1.0 {
			jumpPower *= g.player.speedBoost
		}

		g.player.Jump(jumpPower)
		g.player.charge = 0
	}

	// Player physics (with slow motion effect)
	if g.player.isJumping {
		timeMultiplier := 1.0
		if g.player.slowMotion > 0 && g.player.slowMotion < 1.0 {
			timeMultiplier = g.player.slowMotion
		}

		g.player.vy += gravity * timeMultiplier
		g.player.x += g.player.vx * timeMultiplier
		g.player.y += g.player.vy * timeMultiplier

		// Wall bouncing
		if g.player.x < 0 || g.player.x > screenWidth-playerWidth {
			g.player.vx *= -1
			g.player.direction *= -1
		}
	}

	// Check for power-up collection
	playerRect := g.player.GetRect()
	for _, powerUp := range g.powerUps {
		if !powerUp.collected {
			powerUpRect := powerUp.GetRect()
			if playerRect.Max.X > powerUpRect.Min.X &&
				playerRect.Min.X < powerUpRect.Max.X &&
				playerRect.Max.Y > powerUpRect.Min.Y &&
				playerRect.Min.Y < powerUpRect.Max.Y {
				powerUp.collected = true
				g.applyPowerUp(powerUp.kind)
				g.createParticles(powerUp.x+10, powerUp.y+10, powerUpColors[powerUp.kind])
			}
		}
	}

	// Update power-up timers
	g.updatePowerUps()

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
				if (g.player.y+playerHeight)-g.player.vy <= p.y {
					g.player.isJumping = false
					g.player.vy = 0
					g.player.vx = 0
					g.player.y = p.y - playerHeight

					g.handleScoring(p)
					g.createParticles(g.player.x+playerWidth/2, g.player.y+playerHeight, color.RGBA{255, 255, 255, 255})

					// Generate new platforms if needed
					if p == g.platforms[len(g.platforms)-5] {
						g.generatePlatforms()
					}
					break
				}
			}
		}
	}

	// Update particles
	g.updateParticles()

	// Game over condition
	if g.player.y > g.cameraY+screenHeight {
		g.gameOver = true
		g.saveHighScore()
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

	baseScore := 0
	if dist < platformWidth/4 { // Landed near the center
		g.combo++
		baseScore = g.combo * 2
	} else {
		g.combo = 1 // Reset combo to 1 for a normal hit
		baseScore = 1
	}

	if score, ok := platformScores[p.kind]; ok {
		baseScore += score
	}

	// Apply double score power-up
	if g.player.doubleScore > 0 {
		baseScore *= 2
	}

	g.score += baseScore

	// Update high score
	if g.score > g.highScore {
		g.highScore = g.score
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

	// Draw power-ups
	for _, powerUp := range g.powerUps {
		if !powerUp.collected && powerUp.y > g.cameraY-50 && powerUp.y < g.cameraY+screenHeight+50 {
			powerUpImg := ebiten.NewImage(20, 20)
			powerUpImg.Fill(powerUpColors[powerUp.kind])
			powerUpOp := &ebiten.DrawImageOptions{}
			powerUpOp.GeoM.Translate(powerUp.x, powerUp.y-g.cameraY)
			screen.DrawImage(powerUpImg, powerUpOp)
		}
	}

	// Draw particles
	for _, particle := range g.particles {
		if particle.life > 0 {
			particleImg := ebiten.NewImage(4, 4)
			particleImg.Fill(particle.color)
			particleOp := &ebiten.DrawImageOptions{}
			particleOp.GeoM.Translate(particle.x, particle.y-g.cameraY)
			screen.DrawImage(particleImg, particleOp)
		}
	}

	// Draw player
	playerOp := &ebiten.DrawImageOptions{}
	if g.player.direction == -1 {
		playerOp.GeoM.Scale(-1, 1)
		playerOp.GeoM.Translate(playerWidth, 0)
	}
	playerOp.GeoM.Translate(g.player.x, g.player.y-g.cameraY)
	screen.DrawImage(playerImage, playerOp)

	// Draw UI
	scoreStr := fmt.Sprintf("Score: %d", g.score)
	text.Draw(screen, scoreStr, basicfont.Face7x13, 10, 20, color.Black)

	highScoreStr := fmt.Sprintf("High: %d", g.highScore)
	text.Draw(screen, highScoreStr, basicfont.Face7x13, 10, 40, color.Black)

	if g.isCharging {
		chargeStr := fmt.Sprintf("Power: %.1f", g.player.charge)
		text.Draw(screen, chargeStr, basicfont.Face7x13, 10, 60, color.Black)
	}

	// Draw active power-ups
	yOffset := 80
	for powerType, timer := range g.player.powerUpTimers {
		if timer > 0 {
			powerUpName := []string{"Speed", "2x Score", "Extra Jump", "Slow Mo"}[powerType]
			powerStr := fmt.Sprintf("%s: %.1fs", powerUpName, timer)
			text.Draw(screen, powerStr, basicfont.Face7x13, 10, yOffset, color.Black)
			yOffset += 20
		}
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

// Power-up methods
func (g *Game) applyPowerUp(powerType int) {
	switch powerType {
	case powerUpSpeedBoost:
		g.player.speedBoost = 1.5
		g.player.powerUpTimers[powerType] = 10.0 // 10 seconds
	case powerUpDoubleScore:
		g.player.doubleScore = 2.0
		g.player.powerUpTimers[powerType] = 15.0 // 15 seconds
	case powerUpExtraJump:
		g.player.extraJumps = 1
		g.player.powerUpTimers[powerType] = 20.0 // 20 seconds
	case powerUpSlowMotion:
		g.player.slowMotion = 0.5
		g.player.powerUpTimers[powerType] = 8.0 // 8 seconds
	}
}

func (g *Game) updatePowerUps() {
	for powerType, timer := range g.player.powerUpTimers {
		if timer > 0 {
			g.player.powerUpTimers[powerType] = timer - 1.0/60.0 // Assume 60 FPS
			if g.player.powerUpTimers[powerType] <= 0 {
				g.removePowerUp(powerType)
			}
		}
	}
}

func (g *Game) removePowerUp(powerType int) {
	switch powerType {
	case powerUpSpeedBoost:
		g.player.speedBoost = 1.0
	case powerUpDoubleScore:
		g.player.doubleScore = 1.0
	case powerUpExtraJump:
		g.player.extraJumps = 0
	case powerUpSlowMotion:
		g.player.slowMotion = 1.0
	}
	delete(g.player.powerUpTimers, powerType)
}

// Particle methods
func (g *Game) createParticles(x, y float64, particleColor color.Color) {
	for i := 0; i < 10; i++ {
		particle := &Particle{
			x:     x + rand.Float64()*20 - 10,
			y:     y + rand.Float64()*20 - 10,
			vx:    (rand.Float64() - 0.5) * 4,
			vy:    (rand.Float64() - 0.5) * 4,
			life:  1.0,
			color: particleColor,
		}
		g.particles = append(g.particles, particle)
	}
}

func (g *Game) updateParticles() {
	for i := len(g.particles) - 1; i >= 0; i-- {
		p := g.particles[i]
		p.x += p.vx
		p.y += p.vy
		p.vy += 0.1 // gravity
		p.life -= 0.02

		if p.life <= 0 {
			g.particles = append(g.particles[:i], g.particles[i+1:]...)
		}
	}
}

// High score persistence
func (g *Game) loadHighScore() {
	// For simplicity, we'll use a basic approach. In a real implementation,
	// you might want to use local storage or a file
	g.highScore = 0
}

func (g *Game) saveHighScore() {
	// For simplicity, this is a placeholder. In a real implementation,
	// you would save to local storage or a file
	if g.score > g.highScore {
		g.highScore = g.score
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Jump Jump Game - Enhanced")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
