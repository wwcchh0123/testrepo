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

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

const (
	screenWidth         = 780
	screenHeight        = 1080
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

// Background elements
type Cloud struct {
	x, y   float64
	speed  float64
	scale  float64
	alpha  uint8
}

type Building struct {
	x, y      float64
	width     float64
	height    float64
	color     color.Color
	windows   []Window
}

type Window struct {
	x, y          float64
	width, height float64
	isLit         bool
}

type Star struct {
	x, y      float64
	twinkle   float64
	brightness uint8
}

type BackgroundLayer struct {
	clouds    []Cloud
	buildings []Building
	stars     []Star
	time      float64
}

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

type Game struct {
	player     *Player
	platforms  []*Platform
	cameraY    float64
	score      int
	combo      int
	gameOver   bool
	isCharging bool
	background *BackgroundLayer
}

func NewGame() *Game {
	g := &Game{}
	g.player = &Player{x: screenWidth/2 - playerWidth/2, y: screenHeight - platformHeight - playerHeight, direction: 1}
	g.platforms = append(g.platforms, &Platform{x: screenWidth/2 - platformWidth/2, y: screenHeight - platformHeight, kind: platformNormal})
	g.generatePlatforms()
	g.background = createBackgroundLayer()
	return g
}

func createBackgroundLayer() *BackgroundLayer {
	bg := &BackgroundLayer{}
	
	// Create clouds
	for range 8 {
		cloud := Cloud{
			x:     rand.Float64() * screenWidth * 2,
			y:     rand.Float64() * screenHeight * 0.6,
			speed: rand.Float64()*0.5 + 0.2,
			scale: rand.Float64()*0.5 + 0.5,
			alpha: uint8(rand.Intn(100) + 100),
		}
		bg.clouds = append(bg.clouds, cloud)
	}
	
	// Create buildings
	buildingX := 0.0
	for buildingX < screenWidth*1.5 {
		buildingWidth := rand.Float64()*80 + 40
		buildingHeight := rand.Float64()*200 + 150
		building := Building{
			x:      buildingX,
			y:      screenHeight - buildingHeight,
			width:  buildingWidth,
			height: buildingHeight,
			color:  randomBuildingColor(),
		}
		
		// Add windows
		for wx := building.x + 8; wx < building.x+building.width-8; wx += 12 {
			for wy := building.y + 20; wy < building.y+building.height-10; wy += 20 {
				window := Window{
					x:      wx,
					y:      wy,
					width:  6,
					height: 8,
					isLit:  rand.Float64() < 0.3,
				}
				building.windows = append(building.windows, window)
			}
		}
		
		bg.buildings = append(bg.buildings, building)
		buildingX += buildingWidth + rand.Float64()*20 + 10
	}
	
	// Create stars
	for range 50 {
		star := Star{
			x:          rand.Float64() * screenWidth,
			y:          rand.Float64() * screenHeight * 0.4,
			twinkle:    rand.Float64() * math.Pi * 2,
			brightness: uint8(rand.Intn(155) + 100),
		}
		bg.stars = append(bg.stars, star)
	}
	
	return bg
}

func randomBuildingColor() color.Color {
	colors := []color.Color{
		color.RGBA{60, 60, 80, 255},   // Dark blue-grey
		color.RGBA{80, 70, 90, 255},   // Purple-grey
		color.RGBA{70, 80, 70, 255},   // Green-grey
		color.RGBA{90, 80, 70, 255},   // Brown-grey
		color.RGBA{50, 60, 70, 255},   // Blue-grey
	}
	return colors[rand.Intn(len(colors))]
}

func (g *Game) updateBackground() {
	g.background.time += 0.016 // Assuming ~60 FPS
	
	// Update clouds
	for i := range g.background.clouds {
		g.background.clouds[i].x -= g.background.clouds[i].speed
		if g.background.clouds[i].x < -100 {
			g.background.clouds[i].x = screenWidth + 50
			g.background.clouds[i].y = rand.Float64() * screenHeight * 0.6
		}
	}
	
	// Update stars twinkling
	for i := range g.background.stars {
		g.background.stars[i].twinkle += 0.05
		g.background.stars[i].brightness = uint8(155 + 100*math.Sin(g.background.stars[i].twinkle))
	}
	
	// Occasionally toggle building windows
	if rand.Float64() < 0.01 {
		for i := range g.background.buildings {
			for j := range g.background.buildings[i].windows {
				if rand.Float64() < 0.1 {
					g.background.buildings[i].windows[j].isLit = !g.background.buildings[i].windows[j].isLit
				}
			}
		}
	}
}

func (g *Game) generatePlatforms() {
	for range 10 {
		lastPlatform := g.platforms[len(g.platforms)-1]
		newX := rand.Float64() * (screenWidth - platformWidth)
		newY := lastPlatform.y - float64(rand.Intn(100)+80) // Increase min distance

		kind := platformNormal
		if rand.Float64() < 0.2 { // 20% chance for a special platform
			kind = rand.Intn(len(platformScores)) + 1
		}

		g.platforms = append(g.platforms, &Platform{x: newX, y: newY, kind: kind})
	}
}

func (g *Game) Update() error {
	// Update background animation
	g.updateBackground()
	
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
			g.player.direction *= -1
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
				if (g.player.y+playerHeight)-g.player.vy <= p.y {
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
	// Draw gradient sky background
	g.drawGradientSky(screen)
	
	// Draw background elements with parallax
	g.drawBackground(screen)

	// Draw platforms
	for _, p := range g.platforms {
		platformOp := &ebiten.DrawImageOptions{}
		platformOp.GeoM.Translate(p.x, p.y-g.cameraY)
		platformOp.ColorScale.ScaleWithColor(platformColors[p.kind])
		screen.DrawImage(platformImage, platformOp)
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
	text.Draw(screen, scoreStr, basicfont.Face7x13, 10, 20, color.White)

	if g.isCharging {
		chargeStr := fmt.Sprintf("Power: %.1f", g.player.charge)
		text.Draw(screen, chargeStr, basicfont.Face7x13, 10, 40, color.White)
	}

	if g.gameOver {
		msg := "GAME OVER"
		subMsg := "Press 'R' to Restart"
		msgX := (screenWidth - len(msg)*7) / 2
		subMsgX := (screenWidth - len(subMsg)*7) / 2
		text.Draw(screen, msg, basicfont.Face7x13, msgX, screenHeight/2-20, color.White)
		text.Draw(screen, subMsg, basicfont.Face7x13, subMsgX, screenHeight/2, color.White)
	}
}

func (g *Game) drawGradientSky(screen *ebiten.Image) {
	// Time-based color changes for day/night cycle
	timeOfDay := math.Sin(g.background.time * 0.1)
	
	// Colors for different times
	var topColor, bottomColor color.RGBA
	
	if timeOfDay > 0 { // Day time
		intensity := uint8(timeOfDay * 100 + 100)
		topColor = color.RGBA{135 + intensity/3, 206 + intensity/5, 235 + intensity/4, 255}    // Light blue
		bottomColor = color.RGBA{255 - intensity/2, 240 - intensity/3, 200 - intensity/4, 255} // Light yellow
	} else { // Night time
		intensity := uint8(-timeOfDay * 80 + 50)
		topColor = color.RGBA{25 + intensity/4, 25 + intensity/4, 45 + intensity/2, 255}   // Dark blue
		bottomColor = color.RGBA{60 + intensity/2, 40 + intensity/3, 80 + intensity/2, 255} // Purple
	}
	
	// Draw gradient
	for y := 0; y < screenHeight; y++ {
		ratio := float64(y) / float64(screenHeight)
		r := uint8(float64(topColor.R)*(1-ratio) + float64(bottomColor.R)*ratio)
		g := uint8(float64(topColor.G)*(1-ratio) + float64(bottomColor.G)*ratio)
		b := uint8(float64(topColor.B)*(1-ratio) + float64(bottomColor.B)*ratio)
		
		lineColor := color.RGBA{r, g, b, 255}
		lineImg := ebiten.NewImage(screenWidth, 1)
		lineImg.Fill(lineColor)
		
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, float64(y))
		screen.DrawImage(lineImg, op)
	}
}

func (g *Game) drawBackground(screen *ebiten.Image) {
	// Draw stars (furthest layer, minimal parallax)
	for _, star := range g.background.stars {
		starY := star.y - g.cameraY*0.1 // Very slow parallax
		if starY > -10 && starY < screenHeight+10 {
			starImg := ebiten.NewImage(2, 2)
			starColor := color.RGBA{255, 255, 255, star.brightness}
			starImg.Fill(starColor)
			
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(star.x, starY)
			screen.DrawImage(starImg, op)
		}
	}
	
	// Draw clouds (middle layer, medium parallax)
	for _, cloud := range g.background.clouds {
		cloudY := cloud.y - g.cameraY*0.3
		if cloudY > -50 && cloudY < screenHeight+50 {
			g.drawCloud(screen, cloud.x, cloudY, cloud.scale, cloud.alpha)
		}
	}
	
	// Draw buildings (closest layer, most parallax)
	for _, building := range g.background.buildings {
		buildingY := building.y - g.cameraY*0.7
		if buildingY < screenHeight+50 {
			g.drawBuilding(screen, building, buildingY)
		}
	}
}

func (g *Game) drawCloud(screen *ebiten.Image, x, y, scale float64, alpha uint8) {
	cloudColor := color.RGBA{255, 255, 255, alpha}
	
	// Simple cloud shape using circles
	circles := []struct{ dx, dy, size float64 }{
		{0, 0, 15 * scale},
		{12 * scale, 0, 12 * scale},
		{-8 * scale, 0, 10 * scale},
		{6 * scale, -6 * scale, 8 * scale},
		{-2 * scale, -4 * scale, 6 * scale},
	}
	
	for _, circle := range circles {
		circleImg := ebiten.NewImage(int(circle.size*2), int(circle.size*2))
		circleImg.Fill(cloudColor)
		
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(x+circle.dx-circle.size, y+circle.dy-circle.size)
		screen.DrawImage(circleImg, op)
	}
}

func (g *Game) drawBuilding(screen *ebiten.Image, building Building, y float64) {
	// Draw building body
	buildingImg := ebiten.NewImage(int(building.width), int(building.height))
	buildingImg.Fill(building.color)
	
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(building.x, y)
	screen.DrawImage(buildingImg, op)
	
	// Draw windows
	for _, window := range building.windows {
		if window.isLit {
			windowImg := ebiten.NewImage(int(window.width), int(window.height))
			windowImg.Fill(color.RGBA{255, 255, 150, 255}) // Warm yellow light
			
			windowOp := &ebiten.DrawImageOptions{}
			windowOp.GeoM.Translate(window.x, window.y-building.y+y)
			screen.DrawImage(windowImg, windowOp)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Jump Jump Game")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
