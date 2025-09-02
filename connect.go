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
	"golang.org/x/image/font/basicfont"
)

const (
	connectScreenWidth  = 800
	connectScreenHeight = 600
	boardWidth         = 16
	boardHeight        = 10
	tileSize           = 32
	boardOffsetX       = (connectScreenWidth - boardWidth*tileSize) / 2
	boardOffsetY       = (connectScreenHeight - boardHeight*tileSize) / 2
	tileTypes          = 8
)

// TileType represents different tile images
type TileType int

const (
	TileEmpty TileType = iota
	TileAnimal1 // Cat
	TileAnimal2 // Dog
	TileAnimal3 // Panda
	TileAnimal4 // Tiger
	TileFruit1  // Apple
	TileFruit2  // Orange
	TileFruit3  // Strawberry
	TileFruit4  // Grape
)

var tileColors = map[TileType]color.Color{
	TileEmpty:   color.RGBA{50, 50, 50, 255},
	TileAnimal1: color.RGBA{255, 192, 203, 255}, // Pink (Cat)
	TileAnimal2: color.RGBA{139, 69, 19, 255},   // Brown (Dog)
	TileAnimal3: color.RGBA{0, 0, 0, 255},       // Black (Panda)
	TileAnimal4: color.RGBA{255, 140, 0, 255},   // Orange (Tiger)
	TileFruit1:  color.RGBA{255, 0, 0, 255},     // Red (Apple)
	TileFruit2:  color.RGBA{255, 165, 0, 255},   // Orange
	TileFruit3:  color.RGBA{255, 20, 147, 255},  // Deep Pink (Strawberry)
	TileFruit4:  color.RGBA{128, 0, 128, 255},   // Purple (Grape)
}

// Tile represents a single tile on the board
type Tile struct {
	Type     TileType
	X, Y     int
	Selected bool
	Visible  bool
}

// ConnectGame represents the main connect game state
type ConnectGame struct {
	board         [][]*Tile
	selectedTile  *Tile
	score         int
	gameOver      bool
	gameWon       bool
	connectionPath []image.Point
	showPath      bool
	pathTimer     float64
}

// NewConnectGame creates a new connect game instance
func NewConnectGame() *ConnectGame {
	g := &ConnectGame{}
	g.initializeBoard()
	return g
}

func (g *ConnectGame) initializeBoard() {
	g.board = make([][]*Tile, boardHeight+2) // Add border
	for y := 0; y < boardHeight+2; y++ {
		g.board[y] = make([]*Tile, boardWidth+2) // Add border
		for x := 0; x < boardWidth+2; x++ {
			g.board[y][x] = &Tile{
				Type:    TileEmpty,
				X:       x,
				Y:       y,
				Visible: false,
			}
		}
	}

	// Generate pairs of tiles
	tileCount := (boardWidth * boardHeight) / 2
	if tileCount%2 != 0 {
		tileCount-- // Ensure even number
	}

	tiles := make([]TileType, 0, tileCount*2)
	for i := 0; i < tileCount; i++ {
		tileType := TileType((i % tileTypes) + 1)
		tiles = append(tiles, tileType, tileType) // Add pair
	}

	// Shuffle tiles
	rand.Shuffle(len(tiles), func(i, j int) {
		tiles[i], tiles[j] = tiles[j], tiles[i]
	})

	// Place tiles on board (excluding border)
	tileIndex := 0
	for y := 1; y < boardHeight+1; y++ {
		for x := 1; x < boardWidth+1; x++ {
			if tileIndex < len(tiles) {
				g.board[y][x].Type = tiles[tileIndex]
				g.board[y][x].Visible = true
				tileIndex++
			}
		}
	}
}

// findPath finds a valid path between two tiles
func (g *ConnectGame) findPath(from, to *Tile) []image.Point {
	if from.Type != to.Type || from == to {
		return nil
	}

	// Try direct horizontal line
	if from.Y == to.Y {
		path := g.checkHorizontalPath(from.X, to.X, from.Y)
		if path != nil {
			return path
		}
	}

	// Try direct vertical line
	if from.X == to.X {
		path := g.checkVerticalPath(from.Y, to.Y, from.X)
		if path != nil {
			return path
		}
	}

	// Try one-corner paths
	path := g.checkOneCornerPath(from, to)
	if path != nil {
		return path
	}

	// Try two-corner paths
	return g.checkTwoCornerPath(from, to)
}

func (g *ConnectGame) checkHorizontalPath(x1, x2, y int) []image.Point {
	minX, maxX := x1, x2
	if x1 > x2 {
		minX, maxX = x2, x1
	}

	for x := minX; x <= maxX; x++ {
		if g.board[y][x].Type != TileEmpty && (x != x1 && x != x2) {
			return nil
		}
	}

	path := make([]image.Point, 0)
	for x := minX; x <= maxX; x++ {
		path = append(path, image.Point{x, y})
	}
	return path
}

func (g *ConnectGame) checkVerticalPath(y1, y2, x int) []image.Point {
	minY, maxY := y1, y2
	if y1 > y2 {
		minY, maxY = y2, y1
	}

	for y := minY; y <= maxY; y++ {
		if g.board[y][x].Type != TileEmpty && (y != y1 && y != y2) {
			return nil
		}
	}

	path := make([]image.Point, 0)
	for y := minY; y <= maxY; y++ {
		path = append(path, image.Point{x, y})
	}
	return path
}

func (g *ConnectGame) checkOneCornerPath(from, to *Tile) []image.Point {
	// Try corner at (from.X, to.Y)
	corner1 := image.Point{from.X, to.Y}
	if g.board[corner1.Y][corner1.X].Type == TileEmpty {
		path1 := g.checkVerticalPath(from.Y, to.Y, from.X)
		path2 := g.checkHorizontalPath(from.X, to.X, to.Y)
		if path1 != nil && path2 != nil {
			// Combine paths
			combined := append(path1, path2[1:]...)
			return combined
		}
	}

	// Try corner at (to.X, from.Y)
	corner2 := image.Point{to.X, from.Y}
	if g.board[corner2.Y][corner2.X].Type == TileEmpty {
		path1 := g.checkHorizontalPath(from.X, to.X, from.Y)
		path2 := g.checkVerticalPath(from.Y, to.Y, to.X)
		if path1 != nil && path2 != nil {
			// Combine paths
			combined := append(path1, path2[1:]...)
			return combined
		}
	}

	return nil
}

func (g *ConnectGame) checkTwoCornerPath(from, to *Tile) []image.Point {
	// This is a simplified version - full implementation would be more complex
	return nil
}

func (g *ConnectGame) removeTiles(tile1, tile2 *Tile) {
	tile1.Type = TileEmpty
	tile1.Visible = false
	tile2.Type = TileEmpty  
	tile2.Visible = false
	g.score += 10

	// Check if game is won
	g.checkGameWon()
}

func (g *ConnectGame) checkGameWon() {
	for y := 1; y < boardHeight+1; y++ {
		for x := 1; x < boardWidth+1; x++ {
			if g.board[y][x].Visible {
				return
			}
		}
	}
	g.gameWon = true
}

func (g *ConnectGame) Update() error {
	if g.gameOver || g.gameWon {
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			// Restart game
			g.initializeBoard()
			g.selectedTile = nil
			g.score = 0
			g.gameOver = false
			g.gameWon = false
			g.showPath = false
		}
		return nil
	}

	// Handle path display timer
	if g.showPath {
		g.pathTimer -= 1.0/60.0 // Assume 60 FPS
		if g.pathTimer <= 0 {
			g.showPath = false
		}
	}

	// Handle mouse input
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		
		// Convert screen coordinates to board coordinates
		boardX := (mx - boardOffsetX) / tileSize
		boardY := (my - boardOffsetY) / tileSize
		
		// Add 1 to account for border
		boardX++
		boardY++
		
		if boardX >= 1 && boardX < boardWidth+1 && boardY >= 1 && boardY < boardHeight+1 {
			clickedTile := g.board[boardY][boardX]
			
			if clickedTile.Visible {
				if g.selectedTile == nil {
					// First selection
					g.selectedTile = clickedTile
					clickedTile.Selected = true
				} else if g.selectedTile == clickedTile {
					// Deselect same tile
					g.selectedTile.Selected = false
					g.selectedTile = nil
				} else {
					// Second selection - try to connect
					path := g.findPath(g.selectedTile, clickedTile)
					if path != nil {
						// Valid connection found
						g.connectionPath = path
						g.showPath = true
						g.pathTimer = 1.0 // Show path for 1 second
						g.removeTiles(g.selectedTile, clickedTile)
						g.selectedTile.Selected = false
						g.selectedTile = nil
					} else {
						// Invalid connection - select new tile
						g.selectedTile.Selected = false
						g.selectedTile = clickedTile
						clickedTile.Selected = true
					}
				}
			}
		}
	}

	return nil
}

func (g *ConnectGame) Draw(screen *ebiten.Image) {
	// Clear screen
	screen.Fill(color.RGBA{200, 200, 255, 255})

	// Draw board background
	boardBg := ebiten.NewImage(boardWidth*tileSize, boardHeight*tileSize)
	boardBg.Fill(color.RGBA{100, 100, 100, 100})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(boardOffsetX), float64(boardOffsetY))
	screen.DrawImage(boardBg, op)

	// Draw tiles
	for y := 1; y < boardHeight+1; y++ {
		for x := 1; x < boardWidth+1; x++ {
			tile := g.board[y][x]
			if tile.Visible {
				g.drawTile(screen, tile, x-1, y-1) // Subtract 1 for display coordinates
			}
		}
	}

	// Draw connection path
	if g.showPath && len(g.connectionPath) > 1 {
		g.drawPath(screen)
	}

	// Draw UI
	scoreStr := fmt.Sprintf("Score: %d", g.score)
	text.Draw(screen, scoreStr, basicfont.Face7x13, 10, 20, color.Black)

	if g.gameWon {
		msg := "YOU WIN!"
		subMsg := "Press 'R' to Restart"
		msgX := (connectScreenWidth - len(msg)*7) / 2
		subMsgX := (connectScreenWidth - len(subMsg)*7) / 2
		text.Draw(screen, msg, basicfont.Face7x13, msgX, connectScreenHeight/2-20, color.Green)
		text.Draw(screen, subMsg, basicfont.Face7x13, subMsgX, connectScreenHeight/2, color.Black)
	} else if g.gameOver {
		msg := "GAME OVER"
		subMsg := "Press 'R' to Restart"
		msgX := (connectScreenWidth - len(msg)*7) / 2
		subMsgX := (connectScreenWidth - len(subMsg)*7) / 2
		text.Draw(screen, msg, basicfont.Face7x13, msgX, connectScreenHeight/2-20, color.Red)
		text.Draw(screen, subMsg, basicfont.Face7x13, subMsgX, connectScreenHeight/2, color.Black)
	}
}

func (g *ConnectGame) drawTile(screen *ebiten.Image, tile *Tile, x, y int) {
	tileImg := ebiten.NewImage(tileSize-2, tileSize-2)
	tileColor := tileColors[tile.Type]
	
	if tile.Selected {
		// Add selection highlight
		tileColor = color.RGBA{255, 255, 255, 255}
	}
	
	tileImg.Fill(tileColor)
	
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(boardOffsetX+x*tileSize+1), float64(boardOffsetY+y*tileSize+1))
	screen.DrawImage(tileImg, op)
	
	// Draw tile type indicator (simple text for now)
	var typeChar string
	switch tile.Type {
	case TileAnimal1:
		typeChar = "🐱"
	case TileAnimal2:
		typeChar = "🐶"
	case TileAnimal3:
		typeChar = "🐼"
	case TileAnimal4:
		typeChar = "🐯"
	case TileFruit1:
		typeChar = "🍎"
	case TileFruit2:
		typeChar = "🍊"
	case TileFruit3:
		typeChar = "🍓"
	case TileFruit4:
		typeChar = "🍇"
	}
	
	if typeChar != "" {
		textX := boardOffsetX + x*tileSize + 8
		textY := boardOffsetY + y*tileSize + 20
		text.Draw(screen, typeChar, basicfont.Face7x13, textX, textY, color.White)
	}
}

func (g *ConnectGame) drawPath(screen *ebiten.Image) {
	if len(g.connectionPath) < 2 {
		return
	}

	// Draw lines connecting the path points
	for i := 0; i < len(g.connectionPath)-1; i++ {
		p1 := g.connectionPath[i]
		p2 := g.connectionPath[i+1]
		
		// Convert board coordinates to screen coordinates
		x1 := boardOffsetX + (p1.X-1)*tileSize + tileSize/2
		y1 := boardOffsetY + (p1.Y-1)*tileSize + tileSize/2
		x2 := boardOffsetX + (p2.X-1)*tileSize + tileSize/2
		y2 := boardOffsetY + (p2.Y-1)*tileSize + tileSize/2
		
		g.drawLine(screen, x1, y1, x2, y2, color.RGBA{255, 255, 0, 255})
	}
}

func (g *ConnectGame) drawLine(screen *ebiten.Image, x1, y1, x2, y2 int, c color.Color) {
	// Simple line drawing - draw a thick line using rectangles
	if x1 == x2 {
		// Vertical line
		minY, maxY := y1, y2
		if y1 > y2 {
			minY, maxY = y2, y1
		}
		lineImg := ebiten.NewImage(4, maxY-minY)
		lineImg.Fill(c)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x1-2), float64(minY))
		screen.DrawImage(lineImg, op)
	} else if y1 == y2 {
		// Horizontal line
		minX, maxX := x1, x2
		if x1 > x2 {
			minX, maxX = x2, x1
		}
		lineImg := ebiten.NewImage(maxX-minX, 4)
		lineImg.Fill(c)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(minX), float64(y1-2))
		screen.DrawImage(lineImg, op)
	}
}

func (g *ConnectGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return connectScreenWidth, connectScreenHeight
}

// ConnectMain runs the connect game
func ConnectMain() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(connectScreenWidth, connectScreenHeight)
	ebiten.SetWindowTitle("Connect Game - 连连看")
	if err := ebiten.RunGame(NewConnectGame()); err != nil {
		log.Fatal(err)
	}
}