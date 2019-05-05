package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 800
	screenHeight = 600
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
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}
