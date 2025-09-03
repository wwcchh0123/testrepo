package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

const (
	screenWidth  = 800
	screenHeight = 600
	boardWidth   = 10
	boardHeight  = 8
	tileSize     = 48
	boardOffsetX = (screenWidth - boardWidth*tileSize) / 2
	boardOffsetY = 80

	// Game modes
	modeClassic = iota
	modeTimeAttack
	modeInfinite

	// Tile types (themes)
	themeCount   = 6
	tilesPerType = 4 // 4 tiles of each type for matching
)

// Colors for different tile themes
var tileThemes = [][]color.Color{
	// Fruit theme
	{color.RGBA{255, 100, 100, 255}, color.RGBA{255, 150, 50, 255}, color.RGBA{50, 255, 50, 255}, color.RGBA{150, 50, 255, 255}},
	// Animal theme
	{color.RGBA{139, 69, 19, 255}, color.RGBA{255, 192, 203, 255}, color.RGBA{128, 128, 128, 255}, color.RGBA{0, 0, 139, 255}},
	// Flower theme
	{color.RGBA{255, 20, 147, 255}, color.RGBA{255, 105, 180, 255}, color.RGBA{255, 255, 0, 255}, color.RGBA{138, 43, 226, 255}},
	// Music theme
	{color.RGBA{30, 144, 255, 255}, color.RGBA{255, 215, 0, 255}, color.RGBA{50, 205, 50, 255}, color.RGBA{220, 20, 60, 255}},
	// Food theme
	{color.RGBA{255, 140, 0, 255}, color.RGBA{255, 69, 0, 255}, color.RGBA{34, 139, 34, 255}, color.RGBA{218, 165, 32, 255}},
	// Sport theme
	{color.RGBA{255, 255, 255, 255}, color.RGBA{255, 165, 0, 255}, color.RGBA{0, 128, 0, 255}, color.RGBA{0, 0, 255, 255}},
}

var themeNames = []string{"水果", "动物", "花卉", "音乐", "美食", "运动"}

type Position struct {
	X, Y int
}

type Tile struct {
	TileType int
	Theme    int
	Position Position
	IsEmpty  bool
	Selected bool
	Marked   bool
}

type GameState int

const (
	stateMenu GameState = iota
	statePlaying
	statePaused
	stateGameOver
	stateWin
)

type Game struct {
	board         [][]*Tile
	gameState     GameState
	gameMode      int
	currentTheme  int
	score         int
	timeLeft      int
	combo         int
	selectedTiles []*Tile
	pathPoints    []Position
	showPath      bool
	pathTimer     int

	// Power-ups
	hintsRemaining    int
	shufflesRemaining int
	timeBonus         int

	// UI state
	mouseX, mouseY int
	lastClickTime  time.Time

	// Images
	tileImages []*ebiten.Image
	bgImage    *ebiten.Image

	// Game timer
	gameTimer int
}

func NewGame() *Game {
	g := &Game{
		gameState:         stateMenu,
		gameMode:          modeClassic,
		currentTheme:      0,
		timeLeft:          300, // 5 minutes for classic mode
		hintsRemaining:    3,
		shufflesRemaining: 2,
		gameTimer:         0,
	}

	g.initializeImages()
	g.initializeBoard()

	return g
}

func (g *Game) initializeImages() {
	// Create tile images for different themes
	g.tileImages = make([]*ebiten.Image, themeCount*tilesPerType)

	for theme := 0; theme < themeCount; theme++ {
		for tileType := 0; tileType < tilesPerType; tileType++ {
			img := ebiten.NewImage(tileSize-2, tileSize-2)

			// Fill with theme color
			img.Fill(tileThemes[theme][tileType])

			// Add border
			vector.StrokeRect(img, 1, 1, float32(tileSize-4), float32(tileSize-4), 2, color.RGBA{50, 50, 50, 255}, true)

			// Add simple pattern/symbol in center
			centerX, centerY := float32(tileSize/2), float32(tileSize/2)
			symbolColor := color.RGBA{255, 255, 255, 200}

			switch tileType {
			case 0: // Circle
				vector.DrawFilledCircle(img, centerX, centerY, 8, symbolColor, true)
			case 1: // Square
				vector.DrawFilledRect(img, centerX-6, centerY-6, 12, 12, symbolColor, true)
			case 2: // Triangle
				vector.DrawFilledRect(img, centerX-1, centerY-8, 2, 16, symbolColor, true)
				vector.DrawFilledRect(img, centerX-8, centerY+4, 16, 2, symbolColor, true)
			case 3: // Diamond
				vector.DrawFilledRect(img, centerX-1, centerY-10, 2, 20, symbolColor, true)
				vector.DrawFilledRect(img, centerX-10, centerY-1, 20, 2, symbolColor, true)
			}

			g.tileImages[theme*tilesPerType+tileType] = img
		}
	}

	// Create background
	g.bgImage = ebiten.NewImage(screenWidth, screenHeight)
	g.bgImage.Fill(color.RGBA{240, 248, 255, 255}) // Light blue background
}

func (g *Game) initializeBoard() {
	g.board = make([][]*Tile, boardHeight)
	for y := 0; y < boardHeight; y++ {
		g.board[y] = make([]*Tile, boardWidth)
		for x := 0; x < boardWidth; x++ {
			g.board[y][x] = &Tile{
				Position: Position{X: x, Y: y},
				IsEmpty:  true,
			}
		}
	}

	// Fill board with tiles
	g.generateTiles()
}

func (g *Game) generateTiles() {
	// Calculate how many pairs we need
	totalTiles := boardWidth * boardHeight
	if totalTiles%2 != 0 {
		totalTiles-- // Ensure even number
	}

	pairs := totalTiles / 2
	tilesPerType := pairs / tilesPerType

	// Create pairs of tiles
	var tiles []*Tile
	tileID := 0

	for tileType := 0; tileType < tilesPerType; tileType++ {
		for pair := 0; pair < tilesPerType; pair++ {
			// Create a pair of matching tiles
			for i := 0; i < 2; i++ {
				tiles = append(tiles, &Tile{
					TileType: tileType,
					Theme:    g.currentTheme,
					IsEmpty:  false,
				})
			}
			tileID++
		}
	}

	// Shuffle tiles
	rand.Shuffle(len(tiles), func(i, j int) {
		tiles[i], tiles[j] = tiles[j], tiles[i]
	})

	// Place tiles on board
	tileIndex := 0
	for y := 0; y < boardHeight && tileIndex < len(tiles); y++ {
		for x := 0; x < boardWidth && tileIndex < len(tiles); x++ {
			if tileIndex < len(tiles) {
				tile := tiles[tileIndex]
				tile.Position = Position{X: x, Y: y}
				g.board[y][x] = tile
				tileIndex++
			}
		}
	}

	// Ensure the board is solvable by checking if there's at least one valid path
	if !g.hasSolution() {
		// Regenerate if no solution exists
		g.generateTiles()
	}
}

func (g *Game) hasSolution() bool {
	// Simple check: count matching pairs that can be connected
	matchingPairs := 0

	for y1 := 0; y1 < boardHeight; y1++ {
		for x1 := 0; x1 < boardWidth; x1++ {
			tile1 := g.board[y1][x1]
			if tile1.IsEmpty {
				continue
			}

			for y2 := y1; y2 < boardHeight; y2++ {
				startX := 0
				if y2 == y1 {
					startX = x1 + 1
				}

				for x2 := startX; x2 < boardWidth; x2++ {
					tile2 := g.board[y2][x2]
					if tile2.IsEmpty {
						continue
					}

					if g.tilesMatch(tile1, tile2) && g.canConnect(tile1.Position, tile2.Position) {
						matchingPairs++
					}
				}
			}
		}
	}

	return matchingPairs > 0
}

func (g *Game) Update() error {
	g.gameTimer++

	// Get mouse position
	g.mouseX, g.mouseY = ebiten.CursorPosition()

	switch g.gameState {
	case stateMenu:
		g.updateMenu()
	case statePlaying:
		g.updateGame()
	case stateGameOver, stateWin:
		g.updateGameOver()
	}

	return nil
}

func (g *Game) updateMenu() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.gameState = statePlaying
		g.resetGame()
	}

	// Theme selection
	if inpututil.IsKeyJustPressed(ebiten.KeyT) {
		g.currentTheme = (g.currentTheme + 1) % themeCount
		g.initializeImages()
		g.initializeBoard()
	}

	// Mode selection
	if inpututil.IsKeyJustPressed(ebiten.KeyM) {
		g.gameMode = (g.gameMode + 1) % 3
		switch g.gameMode {
		case modeClassic:
			g.timeLeft = 300 // 5 minutes
		case modeTimeAttack:
			g.timeLeft = 180 // 3 minutes
		case modeInfinite:
			g.timeLeft = 0 // No time limit
		}
	}
}

func (g *Game) updateGame() {
	// Update timer
	if g.gameMode != modeInfinite && g.timeLeft > 0 {
		if g.gameTimer%60 == 0 { // Update every second
			g.timeLeft--
			if g.timeLeft <= 0 {
				g.gameState = stateGameOver
				return
			}
		}
	}

	// Handle mouse clicks
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.handleClick(g.mouseX, g.mouseY)
	}

	// Handle keyboard shortcuts
	if inpututil.IsKeyJustPressed(ebiten.KeyH) && g.hintsRemaining > 0 {
		g.useHint()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) && g.shufflesRemaining > 0 {
		g.shuffle()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		g.gameState = statePaused
	}

	// Update path display timer
	if g.showPath && g.pathTimer > 0 {
		g.pathTimer--
		if g.pathTimer <= 0 {
			g.showPath = false
		}
	}

	// Check win condition
	if g.boardIsEmpty() {
		g.gameState = stateWin
	}
}

func (g *Game) updateGameOver() {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.gameState = stateMenu
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.resetGame()
		g.gameState = statePlaying
	}
}

func (g *Game) resetGame() {
	g.score = 0
	g.combo = 0
	g.selectedTiles = nil
	g.pathPoints = nil
	g.showPath = false
	g.hintsRemaining = 3
	g.shufflesRemaining = 2

	switch g.gameMode {
	case modeClassic:
		g.timeLeft = 300
	case modeTimeAttack:
		g.timeLeft = 180
	case modeInfinite:
		g.timeLeft = 0
	}

	g.initializeBoard()
}

func (g *Game) handleClick(x, y int) {
	// Convert screen coordinates to board coordinates
	boardX := (x - boardOffsetX) / tileSize
	boardY := (y - boardOffsetY) / tileSize

	if boardX < 0 || boardX >= boardWidth || boardY < 0 || boardY >= boardHeight {
		return
	}

	tile := g.board[boardY][boardX]
	if tile.IsEmpty {
		return
	}

	// Handle tile selection
	if len(g.selectedTiles) == 0 {
		// First tile selection
		tile.Selected = true
		g.selectedTiles = []*Tile{tile}
	} else if len(g.selectedTiles) == 1 {
		if g.selectedTiles[0] == tile {
			// Deselect same tile
			tile.Selected = false
			g.selectedTiles = nil
		} else if g.tilesMatch(g.selectedTiles[0], tile) {
			// Check if tiles can be connected
			if g.canConnect(g.selectedTiles[0].Position, tile.Position) {
				// Match found!
				g.removeTiles(g.selectedTiles[0], tile)
				g.updateScore(true)
			} else {
				// Can't connect, show feedback
				g.selectedTiles[0].Selected = false
				tile.Selected = true
				g.selectedTiles = []*Tile{tile}
			}
		} else {
			// Different tile type, switch selection
			g.selectedTiles[0].Selected = false
			tile.Selected = true
			g.selectedTiles = []*Tile{tile}
		}
	}
}

func (g *Game) tilesMatch(tile1, tile2 *Tile) bool {
	return !tile1.IsEmpty && !tile2.IsEmpty && tile1.TileType == tile2.TileType && tile1.Theme == tile2.Theme
}

func (g *Game) canConnect(pos1, pos2 Position) bool {
	// Implement L-shaped path finding (max 2 turns)
	// Try direct horizontal/vertical connection
	if g.hasDirectPath(pos1, pos2) {
		g.pathPoints = []Position{pos1, pos2}
		return true
	}

	// Try one-turn connection
	if g.hasOneTurnPath(pos1, pos2) {
		return true
	}

	// Try two-turn connection
	if g.hasTwoTurnPath(pos1, pos2) {
		return true
	}

	return false
}

func (g *Game) hasDirectPath(pos1, pos2 Position) bool {
	// Check horizontal path
	if pos1.Y == pos2.Y {
		minX, maxX := pos1.X, pos2.X
		if minX > maxX {
			minX, maxX = maxX, minX
		}
		for x := minX + 1; x < maxX; x++ {
			if !g.board[pos1.Y][x].IsEmpty {
				return false
			}
		}
		return true
	}

	// Check vertical path
	if pos1.X == pos2.X {
		minY, maxY := pos1.Y, pos2.Y
		if minY > maxY {
			minY, maxY = maxY, minY
		}
		for y := minY + 1; y < maxY; y++ {
			if !g.board[y][pos1.X].IsEmpty {
				return false
			}
		}
		return true
	}

	return false
}

func (g *Game) hasOneTurnPath(pos1, pos2 Position) bool {
	// Try corner at (pos1.X, pos2.Y)
	corner := Position{X: pos1.X, Y: pos2.Y}
	if ((corner.X == pos1.X && corner.Y == pos1.Y) || (corner.X == pos2.X && corner.Y == pos2.Y) || g.board[corner.Y][corner.X].IsEmpty) &&
		g.hasDirectPath(pos1, corner) && g.hasDirectPath(corner, pos2) {
		g.pathPoints = []Position{pos1, corner, pos2}
		return true
	}

	// Try corner at (pos2.X, pos1.Y)
	corner = Position{X: pos2.X, Y: pos1.Y}
	if ((corner.X == pos1.X && corner.Y == pos1.Y) || (corner.X == pos2.X && corner.Y == pos2.Y) || g.board[corner.Y][corner.X].IsEmpty) &&
		g.hasDirectPath(pos1, corner) && g.hasDirectPath(corner, pos2) {
		g.pathPoints = []Position{pos1, corner, pos2}
		return true
	}

	return false
}

func (g *Game) hasTwoTurnPath(pos1, pos2 Position) bool {
	// Try all possible intermediate points
	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			pos := Position{X: x, Y: y}
			if !g.board[y][x].IsEmpty && (pos.X != pos1.X || pos.Y != pos1.Y) && (pos.X != pos2.X || pos.Y != pos2.Y) {
				continue
			}

			intermediate := Position{X: x, Y: y}
			if g.hasOneTurnPath(pos1, intermediate) && g.hasOneTurnPath(intermediate, pos2) {
				// Combine paths (simplified)
				g.pathPoints = []Position{pos1, intermediate, pos2}
				return true
			}
		}
	}

	return false
}

func (g *Game) removeTiles(tile1, tile2 *Tile) {
	tile1.IsEmpty = true
	tile1.Selected = false
	tile2.IsEmpty = true
	tile2.Selected = false
	g.selectedTiles = nil

	// Show connection path briefly
	g.showPath = true
	g.pathTimer = 30 // Show for half a second at 60 FPS
}

func (g *Game) updateScore(matched bool) {
	if matched {
		g.combo++
		baseScore := 10
		comboBonus := g.combo * 2
		timeBonus := 0

		if g.gameMode == modeTimeAttack {
			timeBonus = g.timeLeft / 10
		}

		g.score += baseScore + comboBonus + timeBonus
	} else {
		g.combo = 0
	}
}

func (g *Game) useHint() {
	if g.hintsRemaining <= 0 {
		return
	}

	// Find a valid matching pair
	for y1 := 0; y1 < boardHeight; y1++ {
		for x1 := 0; x1 < boardWidth; x1++ {
			tile1 := g.board[y1][x1]
			if tile1.IsEmpty {
				continue
			}

			for y2 := 0; y2 < boardHeight; y2++ {
				for x2 := 0; x2 < boardWidth; x2++ {
					if y1 == y2 && x1 == x2 {
						continue
					}

					tile2 := g.board[y2][x2]
					if tile2.IsEmpty {
						continue
					}

					if g.tilesMatch(tile1, tile2) && g.canConnect(tile1.Position, tile2.Position) {
						// Mark tiles as hint
						tile1.Marked = true
						tile2.Marked = true
						g.hintsRemaining--

						// Clear marks after a few seconds
						go func() {
							time.Sleep(3 * time.Second)
							tile1.Marked = false
							tile2.Marked = false
						}()

						return
					}
				}
			}
		}
	}
}

func (g *Game) shuffle() {
	if g.shufflesRemaining <= 0 {
		return
	}

	// Collect all non-empty tiles
	var tiles []*Tile
	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			if !g.board[y][x].IsEmpty {
				tiles = append(tiles, &Tile{
					TileType: g.board[y][x].TileType,
					Theme:    g.board[y][x].Theme,
					IsEmpty:  false,
				})
				g.board[y][x].IsEmpty = true
			}
		}
	}

	// Shuffle tiles
	rand.Shuffle(len(tiles), func(i, j int) {
		tiles[i], tiles[j] = tiles[j], tiles[i]
	})

	// Place shuffled tiles back
	tileIndex := 0
	for y := 0; y < boardHeight && tileIndex < len(tiles); y++ {
		for x := 0; x < boardWidth && tileIndex < len(tiles); x++ {
			if g.board[y][x].IsEmpty {
				tile := tiles[tileIndex]
				tile.Position = Position{X: x, Y: y}
				g.board[y][x] = tile
				tileIndex++
			}
		}
	}

	g.shufflesRemaining--
	g.selectedTiles = nil // Clear selection
}

func (g *Game) boardIsEmpty() bool {
	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			if !g.board[y][x].IsEmpty {
				return false
			}
		}
	}
	return true
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw background
	screen.DrawImage(g.bgImage, nil)

	switch g.gameState {
	case stateMenu:
		g.drawMenu(screen)
	case statePlaying, statePaused:
		g.drawGame(screen)
	case stateGameOver:
		g.drawGameOver(screen)
	case stateWin:
		g.drawWin(screen)
	}
}

func (g *Game) drawMenu(screen *ebiten.Image) {
	// Title
	title := "连连看小游戏"
	titleX := (screenWidth - len(title)*14) / 2
	text.Draw(screen, title, basicfont.Face7x13, titleX, 100, color.RGBA{0, 0, 139, 255})

	// Theme selection
	themeText := fmt.Sprintf("当前主题: %s (按T切换)", themeNames[g.currentTheme])
	themeX := (screenWidth - len(themeText)*7) / 2
	text.Draw(screen, themeText, basicfont.Face7x13, themeX, 200, color.Black)

	// Mode selection
	modes := []string{"经典模式", "限时挑战", "无限模式"}
	modeText := fmt.Sprintf("游戏模式: %s (按M切换)", modes[g.gameMode])
	modeX := (screenWidth - len(modeText)*7) / 2
	text.Draw(screen, modeText, basicfont.Face7x13, modeX, 230, color.Black)

	// Instructions
	instructions := []string{
		"点击两个相同的图案来消除它们",
		"图案之间的连线最多只能转折2次",
		"按H使用提示，按S洗牌",
		"",
		"按回车键或点击鼠标开始游戏",
	}

	for i, instruction := range instructions {
		if instruction == "" {
			continue
		}
		instrX := (screenWidth - len(instruction)*7) / 2
		text.Draw(screen, instruction, basicfont.Face7x13, instrX, 300+i*25, color.RGBA{50, 50, 50, 255})
	}

	// Draw sample tiles
	g.drawSampleTiles(screen)
}

func (g *Game) drawSampleTiles(screen *ebiten.Image) {
	startX := (screenWidth - tilesPerType*tileSize) / 2
	startY := 450

	for i := 0; i < tilesPerType; i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(startX+i*tileSize), float64(startY))
		screen.DrawImage(g.tileImages[g.currentTheme*tilesPerType+i], op)
	}
}

func (g *Game) drawGame(screen *ebiten.Image) {
	// Draw game board
	g.drawBoard(screen)

	// Draw UI
	g.drawUI(screen)

	// Draw connection path
	if g.showPath && len(g.pathPoints) > 1 {
		g.drawPath(screen)
	}

	if g.gameState == statePaused {
		g.drawPauseOverlay(screen)
	}
}

func (g *Game) drawBoard(screen *ebiten.Image) {
	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			tile := g.board[y][x]
			if tile.IsEmpty {
				continue
			}

			screenX := boardOffsetX + x*tileSize
			screenY := boardOffsetY + y*tileSize

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(screenX), float64(screenY))

			// Apply selection highlight
			if tile.Selected {
				op.ColorM.Scale(1.2, 1.2, 0.8, 1.0) // Yellow tint
			}

			// Apply hint highlight
			if tile.Marked {
				op.ColorM.Scale(0.8, 1.2, 0.8, 1.0) // Green tint
			}

			imageIndex := tile.Theme*tilesPerType + tile.TileType
			if imageIndex < len(g.tileImages) {
				screen.DrawImage(g.tileImages[imageIndex], op)
			}
		}
	}
}

func (g *Game) drawUI(screen *ebiten.Image) {
	// Score
	scoreText := fmt.Sprintf("得分: %d", g.score)
	text.Draw(screen, scoreText, basicfont.Face7x13, 10, 20, color.Black)

	// Combo
	if g.combo > 1 {
		comboText := fmt.Sprintf("连击: %dx", g.combo)
		text.Draw(screen, comboText, basicfont.Face7x13, 10, 40, color.RGBA{255, 100, 0, 255})
	}

	// Time (if not infinite mode)
	if g.gameMode != modeInfinite {
		minutes := g.timeLeft / 60
		seconds := g.timeLeft % 60
		timeText := fmt.Sprintf("时间: %02d:%02d", minutes, seconds)
		timeColor := color.Black
		if g.timeLeft < 30 {
			timeColor = color.RGBA{255, 0, 0, 255} // Red when low
		}
		text.Draw(screen, timeText, basicfont.Face7x13, screenWidth-150, 20, timeColor)
	}

	// Power-ups
	hintText := fmt.Sprintf("提示: %d (按H)", g.hintsRemaining)
	text.Draw(screen, hintText, basicfont.Face7x13, 10, screenHeight-60, color.RGBA{0, 100, 200, 255})

	shuffleText := fmt.Sprintf("洗牌: %d (按S)", g.shufflesRemaining)
	text.Draw(screen, shuffleText, basicfont.Face7x13, 10, screenHeight-40, color.RGBA{0, 100, 200, 255})

	// Controls
	text.Draw(screen, "按P暂停", basicfont.Face7x13, screenWidth-100, screenHeight-40, color.RGBA{100, 100, 100, 255})
}

func (g *Game) drawPath(screen *ebiten.Image) {
	if len(g.pathPoints) < 2 {
		return
	}

	pathColor := color.RGBA{255, 0, 0, 150}

	for i := 0; i < len(g.pathPoints)-1; i++ {
		p1 := g.pathPoints[i]
		p2 := g.pathPoints[i+1]

		x1 := float32(boardOffsetX + p1.X*tileSize + tileSize/2)
		y1 := float32(boardOffsetY + p1.Y*tileSize + tileSize/2)
		x2 := float32(boardOffsetX + p2.X*tileSize + tileSize/2)
		y2 := float32(boardOffsetY + p2.Y*tileSize + tileSize/2)

		vector.StrokeLine(screen, x1, y1, x2, y2, 3, pathColor, true)
	}
}

func (g *Game) drawPauseOverlay(screen *ebiten.Image) {
	// Semi-transparent overlay
	overlay := ebiten.NewImage(screenWidth, screenHeight)
	overlay.Fill(color.RGBA{0, 0, 0, 128})
	screen.DrawImage(overlay, nil)

	// Pause text
	pauseText := "游戏暂停"
	pauseX := (screenWidth - len(pauseText)*14) / 2
	text.Draw(screen, pauseText, basicfont.Face7x13, pauseX, screenHeight/2-20, color.White)

	continueText := "按P继续游戏"
	continueX := (screenWidth - len(continueText)*7) / 2
	text.Draw(screen, continueText, basicfont.Face7x13, continueX, screenHeight/2+20, color.White)
}

func (g *Game) drawGameOver(screen *ebiten.Image) {
	g.drawGame(screen)

	// Game over overlay
	overlay := ebiten.NewImage(screenWidth, screenHeight)
	overlay.Fill(color.RGBA{0, 0, 0, 180})
	screen.DrawImage(overlay, nil)

	// Game over text
	gameOverText := "游戏结束"
	gameOverX := (screenWidth - len(gameOverText)*14) / 2
	text.Draw(screen, gameOverText, basicfont.Face7x13, gameOverX, screenHeight/2-60, color.RGBA{255, 100, 100, 255})

	finalScoreText := fmt.Sprintf("最终得分: %d", g.score)
	scoreX := (screenWidth - len(finalScoreText)*7) / 2
	text.Draw(screen, finalScoreText, basicfont.Face7x13, scoreX, screenHeight/2-20, color.White)

	restartText := "按R返回菜单或回车重新开始"
	restartX := (screenWidth - len(restartText)*7) / 2
	text.Draw(screen, restartText, basicfont.Face7x13, restartX, screenHeight/2+20, color.White)
}

func (g *Game) drawWin(screen *ebiten.Image) {
	g.drawGame(screen)

	// Win overlay
	overlay := ebiten.NewImage(screenWidth, screenHeight)
	overlay.Fill(color.RGBA{0, 100, 0, 180})
	screen.DrawImage(overlay, nil)

	// Win text
	winText := "恭喜过关！"
	winX := (screenWidth - len(winText)*14) / 2
	text.Draw(screen, winText, basicfont.Face7x13, winX, screenHeight/2-60, color.RGBA{255, 255, 100, 255})

	finalScoreText := fmt.Sprintf("最终得分: %d", g.score)
	scoreX := (screenWidth - len(finalScoreText)*7) / 2
	text.Draw(screen, finalScoreText, basicfont.Face7x13, scoreX, screenHeight/2-20, color.White)

	bonusText := ""
	if g.gameMode != modeInfinite {
		bonus := g.timeLeft * 5
		g.score += bonus
		bonusText = fmt.Sprintf("时间奖励: %d", bonus)
	}

	if bonusText != "" {
		bonusX := (screenWidth - len(bonusText)*7) / 2
		text.Draw(screen, bonusText, basicfont.Face7x13, bonusX, screenHeight/2, color.RGBA{100, 255, 100, 255})
	}

	continueText := "按R返回菜单或回车重新开始"
	continueX := (screenWidth - len(continueText)*7) / 2
	text.Draw(screen, continueText, basicfont.Face7x13, continueX, screenHeight/2+40, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("连连看小游戏 - Match-Match Game")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
