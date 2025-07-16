package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	TetrisScreenWidth  = 300
	TetrisScreenHeight = 600
	BlockSize          = 30
	BoardWidth         = 10
	BoardHeight        = 20
)

var (
	// Tetromino shapes
	Shapes = [][][]int{
		{{1, 1, 1, 1}}, // I
		{{1, 1}, {1, 1}}, // O
		{{0, 1, 0}, {1, 1, 1}}, // T
		{{1, 0, 0}, {1, 1, 1}}, // J
		{{0, 0, 1}, {1, 1, 1}}, // L
		{{1, 1, 0}, {0, 1, 1}}, // S
		{{0, 1, 1}, {1, 1, 0}}, // Z
	}

	// Tetromino colors
	Colors = []color.Color{
		color.RGBA{0, 255, 255, 255}, // Cyan
		color.RGBA{255, 255, 0, 255}, // Yellow
		color.RGBA{128, 0, 128, 255}, // Purple
		color.RGBA{0, 0, 255, 255},   // Blue
		color.RGBA{255, 165, 0, 255}, // Orange
		color.RGBA{0, 255, 0, 255},   // Green
		color.RGBA{255, 0, 0, 255},   // Red
	}
)

type Tetromino struct {
	Shape      [][]int
	Color      color.Color
	X, Y       int
	ShapeIndex int
}

func NewTetromino() *Tetromino {
	shapeIndex := rand.Intn(len(Shapes))
	return &Tetromino{
		Shape:      Shapes[shapeIndex],
		Color:      Colors[shapeIndex],
		X:          BoardWidth/2 - 2,
		Y:          0,
		ShapeIndex: shapeIndex,
	}
}

func (t *Tetromino) Rotate() {
	newShape := make([][]int, len(t.Shape[0]))
	for i := range newShape {
		newShape[i] = make([]int, len(t.Shape))
	}
	for i, row := range t.Shape {
		for j, val := range row {
			newShape[j][len(t.Shape)-1-i] = val
		}
	}
	t.Shape = newShape
}

type TetrisGame struct {
	Board         [][]color.Color
	CurrentPiece  *Tetromino
	FallTimer     int
	GameOver      bool
	MoveTimer     int
	RotateTimer   int
}

func NewTetrisGame() *TetrisGame {
	board := make([][]color.Color, BoardHeight)
	for i := range board {
		board[i] = make([]color.Color, BoardWidth)
	}
	return &TetrisGame{
		Board:        board,
		CurrentPiece: NewTetromino(),
	}
}

func (g *TetrisGame) Update() error {
	if g.GameOver {
		return nil
	}

	g.FallTimer++
	g.MoveTimer++
    g.RotateTimer++

	// Handle input
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && g.MoveTimer > 5 {
		if g.isValidMove(g.CurrentPiece, g.CurrentPiece.X-1, g.CurrentPiece.Y) {
			g.CurrentPiece.X--
		}
		g.MoveTimer = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) && g.MoveTimer > 5 {
		if g.isValidMove(g.CurrentPiece, g.CurrentPiece.X+1, g.CurrentPiece.Y) {
			g.CurrentPiece.X++
		}
		g.MoveTimer = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) { // Soft drop
		if g.isValidMove(g.CurrentPiece, g.CurrentPiece.X, g.CurrentPiece.Y+1) {
			g.CurrentPiece.Y++
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) && g.RotateTimer > 10 {
		tempPiece := *g.CurrentPiece
		tempShape := make([][]int, len(g.CurrentPiece.Shape))
        for i := range g.CurrentPiece.Shape {
            tempShape[i] = make([]int, len(g.CurrentPiece.Shape[i]))
            copy(tempShape[i], g.CurrentPiece.Shape[i])
        }
        tempPiece.Shape = tempShape
		tempPiece.Rotate()
		if g.isValidMove(&tempPiece, tempPiece.X, tempPiece.Y) {
			g.CurrentPiece.Rotate()
		}
		g.RotateTimer = 0
	}

	// Automatic fall
	if g.FallTimer > 30 {
		if g.isValidMove(g.CurrentPiece, g.CurrentPiece.X, g.CurrentPiece.Y+1) {
			g.CurrentPiece.Y++
		} else {
			g.placePiece()
			g.clearLines()
			g.CurrentPiece = NewTetromino()
			if !g.isValidMove(g.CurrentPiece, g.CurrentPiece.X, g.CurrentPiece.Y) {
				g.GameOver = true
			}
		}
		g.FallTimer = 0
	}

	return nil
}

func (g *TetrisGame) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	// Draw board
	for y, row := range g.Board {
		for x, c := range row {
			if c != nil {
				ebitenutil.DrawRect(screen, float64(x*BlockSize), float64(y*BlockSize), BlockSize-1, BlockSize-1, c)
			}
		}
	}

	// Draw current piece
	for r, row := range g.CurrentPiece.Shape {
		for c, val := range row {
			if val == 1 {
				ebitenutil.DrawRect(screen, float64((g.CurrentPiece.X+c)*BlockSize), float64((g.CurrentPiece.Y+r)*BlockSize), BlockSize-1, BlockSize-1, g.CurrentPiece.Color)
			}
		}
	}
}

func (g *TetrisGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return TetrisScreenWidth, TetrisScreenHeight
}

func (g *TetrisGame) isValidMove(piece *Tetromino, x, y int) bool {
	for r, row := range piece.Shape {
		for c, val := range row {
			if val == 1 {
				newX := x + c
				newY := y + r
				if newX < 0 || newX >= BoardWidth || newY >= BoardHeight {
					return false
				}
				if newY >= 0 && g.Board[newY][newX] != nil {
					return false
				}
			}
		}
	}
	return true
}

func (g *TetrisGame) placePiece() {
	for r, row := range g.CurrentPiece.Shape {
		for c, val := range row {
			if val == 1 {
                if g.CurrentPiece.Y+r >= 0 {
				    g.Board[g.CurrentPiece.Y+r][g.CurrentPiece.X+c] = g.CurrentPiece.Color
                }
			}
		}
	}
}

func (g *TetrisGame) clearLines() {
	newBoard := make([][]color.Color, BoardHeight)
	for i := range newBoard {
		newBoard[i] = make([]color.Color, BoardWidth)
	}

	row := BoardHeight - 1
	for r := BoardHeight - 1; r >= 0; r-- {
		isFull := true
		for c := 0; c < BoardWidth; c++ {
			if g.Board[r][c] == nil {
				isFull = false
				break
			}
		}
		if !isFull {
            if row >= 0 {
			    newBoard[row] = g.Board[r]
			    row--
            }
		}
	}
	g.Board = newBoard
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
