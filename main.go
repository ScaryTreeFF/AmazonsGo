package main

import (
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"fmt"
	"log"
	// "image"
	// "image/color"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
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
	currX, currY := ebiten.CursorPosition()
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		i, j := board.Coor2Ind(currX, currY)
		if board.Pieces[i][j] != nil {
			// make selections
			fmt.Printf("%v %v", i, j)
			x, y := board.Ind2Coor(i, j)
			fmt.Printf("x: %v, y: %v", x, y)
		}
	}
	// is drawing skipped
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// drawing
	board.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("x: %v, y: %v", currX, currY))
	return nil
}

func main() {
	mask := [][]int{
		{0, 0, 1, 2, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{2, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 2},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 2, 1, 0, 0},
	}
	tileNum := 6
	offset := int(screenHeight * 0.1)
	board = NewBoard(offset, offset, tileNum, int(screenHeight * 0.8) / tileNum)
	board.Initialize(mask)
	fmt.Println(board)
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "AmazonsGo"); err != nil {
		log.Fatal(err)
	}
}
