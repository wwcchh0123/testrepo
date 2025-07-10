package main

import (
	"container/list"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
	GridSize     = 10
)

type Game struct {
	snake    *Snake
	food     *Food
	score    int
	gameOver bool
}

type Snake struct {
	body      *list.List
	direction Point
	speed     int
}

type Food struct {
	position Point
}

type Point struct {
	X, Y int
}

func NewGame() *Game {
	g := &Game{}
	g.snake = NewSnake()
	g.food = NewFood()
	g.score = 0
	g.gameOver = false
	return g
}

func NewSnake() *Snake {
	s := &Snake{
		body:      list.New(),
		direction: Point{X: 1, Y: 0},
		speed:     5,
	}
	s.body.PushFront(Point{X: 5, Y: 5})
	return s
}

func NewFood() *Food {
	return &Food{
		position: Point{
			X: rand.Intn(ScreenWidth / GridSize),
			Y: rand.Intn(ScreenHeight / GridSize),
		},
	}
}

func (g *Game) Update() error {
	if g.gameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			g.Reset()
		}
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) && g.snake.direction.Y == 0 {
		g.snake.direction = Point{X: 0, Y: -1}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) && g.snake.direction.Y == 0 {
		g.snake.direction = Point{X: 0, Y: 1}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) && g.snake.direction.X == 0 {
		g.snake.direction = Point{X: -1, Y: 0}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) && g.snake.direction.X == 0 {
		g.snake.direction = Point{X: 1, Y: 0}
	}

	head := g.snake.body.Front().Value.(Point)
	newHead := Point{
		X: head.X + g.snake.direction.X,
		Y: head.Y + g.snake.direction.Y,
	}

	if newHead.X < 0 || newHead.Y < 0 || newHead.X >= ScreenWidth/GridSize || newHead.Y >= ScreenHeight/GridSize {
		g.gameOver = true
	}

	for e := g.snake.body.Front(); e != nil; e = e.Next() {
		if p := e.Value.(Point); p.X == newHead.X && p.Y == newHead.Y {
			g.gameOver = true
			break
		}
	}

	g.snake.body.PushFront(newHead)

	if newHead.X == g.food.position.X && newHead.Y == g.food.position.Y {
		g.score++
		g.food = NewFood()
	} else {
		g.snake.body.Remove(g.snake.body.Back())
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.gameOver {
		ebitenutil.DebugPrint(screen, "Game Over! Press 'R' to Restart")
		return
	}

	for e := g.snake.body.Front(); e != nil; e = e.Next() {
		p := e.Value.(Point)
		ebitenutil.DrawRect(screen, float64(p.X*GridSize), float64(p.Y*GridSize), GridSize, GridSize, color.White)
	}

	ebitenutil.DrawRect(screen, float64(g.food.position.X*GridSize), float64(g.food.position.Y*GridSize), GridSize, GridSize, color.RGBA{R: 255, G: 0, B: 0, A: 255})
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", g.score))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Reset() {
	g.snake = NewSnake()
	g.food = NewFood()
	g.score = 0
	g.gameOver = false
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game := NewGame()
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Snake Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}