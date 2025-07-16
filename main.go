package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	game := NewTetrisGame()
	ebiten.SetWindowSize(TetrisScreenWidth, TetrisScreenHeight)
	ebiten.SetWindowTitle("Tetris")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}