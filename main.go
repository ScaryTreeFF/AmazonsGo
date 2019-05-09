package main

import (
	"fmt"
	"log"
	// "image"
	// "image/color"
	"github.com/hajimehoshi/ebiten"
	// "github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 800
	screenHeight = 800

)

var (
	board Board
)

func update(screen *ebiten.Image) error {
	// updating game state
	// -------------------

	// is drawing skipped
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// drawing
	board.Draw(screen)
	// ebitenutil.DebugPrint(screen, "Hello, World!")
	return nil
}

func main() {
	tileNum := 6
	board = NewBoard(0, 0, tileNum, int(screenHeight * 0.8) / tileNum, 6)
	fmt.Println(board)
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}
