package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	boardWidth       = 10
	boardHeight      = 30
	blockSize        = 30
	boardPixelWidth  = boardWidth * blockSize
	boardPixelHeight = boardHeight * blockSize
	fallSpeed        = 20 // Lower is faster
)

var (
	tetrominoes = [][][]int{
		{{1, 1, 1, 1}}, // I
		{{1, 1}, {1, 1}},   // O
		{{0, 1, 0}, {1, 1, 1}}, // T
		{{0, 1, 1}, {1, 1, 0}}, // S
		{{1, 1, 0}, {0, 1, 1}}, // Z
		{{1, 0, 0}, {1, 1, 1}}, // J
		{{0, 0, 1}, {1, 1, 1}}, // L
	}

	tetrominoColors = []color.Color{
		color.RGBA{0, 255, 255, 255}, // I (Cyan)
		color.RGBA{255, 255, 0, 255}, // O (Yellow)
		color.RGBA{128, 0, 128, 255}, // T (Purple)
		color.RGBA{0, 255, 0, 255},   // S (Green)
		color.RGBA{255, 0, 0, 255},   // Z (Red)
		color.RGBA{0, 0, 255, 255},   // J (Blue)
		color.RGBA{255, 165, 0, 255}, // L (Orange)
	}
	boardColor = color.RGBA{50, 50, 50, 255}
)

type TetrisGame struct {
	board         [boardHeight][boardWidth]int
	currentPiece  [][]int
	pieceX, pieceY int
	pieceColor    color.Color
	fallTimer     int
	gameOver      bool
}

func NewTetrisGame() *TetrisGame {
	tg := &TetrisGame{}
	tg.spawnNewPiece()
	return tg
}

func (tg *TetrisGame) spawnNewPiece() {
	pieceIndex := rand.Intn(len(tetrominoes))
	tg.currentPiece = tetrominoes[pieceIndex]
	tg.pieceColor = tetrominoColors[pieceIndex]
	tg.pieceX = boardWidth/2 - len(tg.currentPiece[0])/2
	tg.pieceY = 0

	if tg.checkCollision(tg.pieceX, tg.pieceY, tg.currentPiece) {
		tg.gameOver = true
	}
}

func (tg *TetrisGame) checkCollision(x, y int, piece [][]int) bool {
	for r, row := range piece {
		for c, cell := range row {
			if cell != 0 {
				boardX := x + c
				boardY := y + r
				if boardX < 0 || boardX >= boardWidth || boardY < 0 || boardY >= boardHeight || tg.board[boardY][boardX] != 0 {
					return true
				}
			}
		}
	}
	return false
}

func (tg *TetrisGame) lockPiece() {
	for r, row := range tg.currentPiece {
		for c, cell := range row {
			if cell != 0 {
				tg.board[tg.pieceY+r][tg.pieceX+c] = 1 // Use 1 to indicate a locked block, could be color index + 1
			}
		}
	}
	tg.clearLines()
	tg.spawnNewPiece()
}

func (tg *TetrisGame) clearLines() {
	for r := boardHeight - 1; r >= 0; r-- {
		fullLine := true
		for c := 0; c < boardWidth; c++ {
			if tg.board[r][c] == 0 {
				fullLine = false
				break
			}
		}
		if fullLine {
			for y := r; y > 0; y-- {
				for x := 0; x < boardWidth; x++ {
					tg.board[y][x] = tg.board[y-1][x]
				}
			}
			for x := 0; x < boardWidth; x++ {
				tg.board[0][x] = 0
			}
			r++ // Re-check the same line
		}
	}
}

func (tg *TetrisGame) rotatePiece() {
	newPiece := make([][]int, len(tg.currentPiece[0]))
	for i := range newPiece {
		newPiece[i] = make([]int, len(tg.currentPiece))
	}
	for r, row := range tg.currentPiece {
		for c, cell := range row {
			newPiece[c][len(tg.currentPiece)-1-r] = cell
		}
	}
	if !tg.checkCollision(tg.pieceX, tg.pieceY, newPiece) {
		tg.currentPiece = newPiece
	}
}

func (tg *TetrisGame) Update() error {
	if tg.gameOver {
		return nil
	}

	// Handle input
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		if !tg.checkCollision(tg.pieceX-1, tg.pieceY, tg.currentPiece) {
			tg.pieceX--
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		if !tg.checkCollision(tg.pieceX+1, tg.pieceY, tg.currentPiece) {
			tg.pieceX++
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		if !tg.checkCollision(tg.pieceX, tg.pieceY+1, tg.currentPiece) {
			tg.pieceY++
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		tg.rotatePiece()
	}

	// Automatic fall
	tg.fallTimer++
	if tg.fallTimer >= fallSpeed {
		tg.fallTimer = 0
		if !tg.checkCollision(tg.pieceX, tg.pieceY+1, tg.currentPiece) {
			tg.pieceY++
		} else {
			tg.lockPiece()
		}
	}

	return nil
}

func (tg *TetrisGame) Draw(screen *ebiten.Image) {
	screen.Fill(boardColor)

	// Draw the board
	for y, row := range tg.board {
		for x, cell := range row {
			if cell != 0 {
				block := ebiten.NewImage(blockSize-1, blockSize-1)
				block.Fill(color.White) // Use a generic color for locked blocks for now
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*blockSize), float64(y*blockSize))
				screen.DrawImage(block, op)
			}
		}
	}

	// Draw the current piece
	block := ebiten.NewImage(blockSize-1, blockSize-1)
	block.Fill(tg.pieceColor)
	op := &ebiten.DrawImageOptions{}
	for y, row := range tg.currentPiece {
		for x, cell := range row {
			if cell != 0 {
				op.GeoM.Reset()
				op.GeoM.Translate(float64((tg.pieceX+x)*blockSize), float64((tg.pieceY+y)*blockSize))
				screen.DrawImage(block, op)
			}
		}
	}
}

func (tg *TetrisGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return boardPixelWidth, boardPixelHeight
}
