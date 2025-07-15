package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

type Game struct {
	currentTime string
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	t := time.Now()
	g.currentTime = t.Format("2006-01-02 15:04:05")
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, g.currentTime)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	game := NewGame()
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Clock")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
