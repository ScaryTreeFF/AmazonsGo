package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

const (
	tileSize  = 10 
)

func update(screen *ebiten.Image) error {
	// updating game state
	// -------------------
	// is drawing skipped
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	// drawing
	ebitenutil.DebugPrint(screen, "Hello, World!")
	return nil
}

func main() {
	board := NewBoard(0, 0, 10, 4)
	fmt.Println(board)
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}
