package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

const (
	screenWidth  = 300
	screenHeight = 600
	boardWidth   = 10
	boardHeight  = 20
	blockSize    = 30
)

var (
	// Tetromino shapes
	tetrominoes = [][][]int{
		{{1, 1, 1, 1}}, // I
		{{1, 1}, {1, 1}}, // O
		{{0, 1, 0}, {1, 1, 1}}, // T
		{{1, 0, 0}, {1, 1, 1}}, // J
		{{0, 0, 1}, {1, 1, 1}}, // L
		{{1, 1, 0}, {0, 1, 1}}, // S
		{{0, 1, 1}, {1, 1, 0}}, // Z
	}

	// Tetromino colors
	tetrominoColors = []color.Color{
		color.RGBA{0, 255, 255, 255}, // Cyan for I
		color.RGBA{255, 255, 0, 255}, // Yellow for O
		color.RGBA{128, 0, 128, 255}, // Purple for T
		color.RGBA{0, 0, 255, 255},   // Blue for J
		color.RGBA{255, 165, 0, 255}, // Orange for L
		color.RGBA{0, 255, 0, 255},   // Green for S
		color.RGBA{255, 0, 0, 255},   // Red for Z
	}
)

type Game struct {
	board         [boardHeight][boardWidth]int
	currentPiece  [][]int
	currentX      int
	currentY      int
	currentColor  color.Color
	score         int
	gameOver      bool
	fallTimer     int
	moveDownTimer int
}

func NewGame() *Game {
	g := &Game{}
	g.spawnNewPiece()
	return g
}

func (g *Game) spawnNewPiece() {
	rand.Seed(time.Now().UnixNano())
	pieceIndex := rand.Intn(len(tetrominoes))
	g.currentPiece = tetrominoes[pieceIndex]
	g.currentColor = tetrominoColors[pieceIndex]
	g.currentX = boardWidth/2 - len(g.currentPiece[0])/2
	g.currentY = 0

	if g.checkCollision(g.currentX, g.currentY, g.currentPiece) {
		g.gameOver = true
	}
}

func (g *Game) checkCollision(x, y int, piece [][]int) bool {
	for r, row := range piece {
		for c, cell := range row {
			if cell != 0 {
				boardX := x + c
				boardY := y + r
				if boardX < 0 || boardX >= boardWidth || boardY >= boardHeight || (boardY >= 0 && g.board[boardY][boardX] != 0) {
					return true
				}
			}
		}
	}
	return false
}

func (g *Game) lockPiece() {
	for r, row := range g.currentPiece {
		for c, cell := range row {
			if cell != 0 {
				if g.currentY+r >= 0 {
					g.board[g.currentY+r][g.currentX+c] = 1
				}
			}
		}
	}
	g.clearLines()
	g.spawnNewPiece()
}

func (g *Game) clearLines() {
	linesCleared := 0
	for r := boardHeight - 1; r >= 0; r-- {
		isFull := true
		for c := 0; c < boardWidth; c++ {
			if g.board[r][c] == 0 {
				isFull = false
				break
			}
		}
		if isFull {
			linesCleared++
			for row := r; row > 0; row-- {
				g.board[row] = g.board[row-1]
			}
			g.board[0] = [boardWidth]int{}
			r++ // Re-check the same line
		}
	}
	g.score += linesCleared * 100
}

func (g *Game) rotatePiece() {
	newPiece := make([][]int, len(g.currentPiece[0]))
	for i := range newPiece {
		newPiece[i] = make([]int, len(g.currentPiece))
	}
	for r, row := range g.currentPiece {
		for c, cell := range row {
			newPiece[c][len(g.currentPiece)-1-r] = cell
		}
	}
	if !g.checkCollision(g.currentX, g.currentY, newPiece) {
		g.currentPiece = newPiece
	}
}

func (g *Game) Update() error {
	if g.gameOver {
		if ebiten.IsKeyPressed(ebiten.KeyR) {
			*g = *NewGame()
		}
		return nil
	}

	// Handle input
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && !g.checkCollision(g.currentX-1, g.currentY, g.currentPiece) {
		g.currentX--
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) && !g.checkCollision(g.currentX+1, g.currentY, g.currentPiece) {
		g.currentX++
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.moveDownTimer++
		if g.moveDownTimer > 5 {
			if !g.checkCollision(g.currentX, g.currentY+1, g.currentPiece) {
				g.currentY++
			} else {
				g.lockPiece()
			}
			g.moveDownTimer = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.rotatePiece()
	}

	// Automatic fall
	g.fallTimer++
	if g.fallTimer > 30 { // Adjust fall speed
		if !g.checkCollision(g.currentX, g.currentY+1, g.currentPiece) {
			g.currentY++
		} else {
			g.lockPiece()
		}
		g.fallTimer = 0
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{50, 50, 50, 255})

	// Draw the board
	for r, row := range g.board {
		for c, cell := range row {
			if cell != 0 {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(c*blockSize), float64(r*blockSize))
				block := ebiten.NewImage(blockSize-1, blockSize-1)
				block.Fill(color.White)
				screen.DrawImage(block, op)
			}
		}
	}

	// Draw the current piece
	op := &ebiten.DrawImageOptions{}
	for r, row := range g.currentPiece {
		for c, cell := range row {
			if cell != 0 {
				op.GeoM.Reset()
				op.GeoM.Translate(float64((g.currentX+c)*blockSize), float64((g.currentY+r)*blockSize))
				block := ebiten.NewImage(blockSize-1, blockSize-1)
				block.Fill(g.currentColor)
				screen.DrawImage(block, op)
			}
		}
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
	game := NewGame()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Tetris")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}