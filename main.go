package main

import (
	"fmt"
	"log"
	// "image"
	"image/color"
	"github.com/hajimehoshi/ebiten"
	// "github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 800
	screenHeight = 800
)

const (
	tileSize  = int(screenHeight / 4)
)

func update(screen *ebiten.Image) error {
	// updating game state
	// -------------------
	// is drawing skipped
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// drawing
	boardImage, _ := ebiten.NewImage(screenWidth, screenHeight, ebiten.FilterDefault)
	boardImage.Fill(color.NRGBA{0x00, 0x40, 0x80, 0xff})
	screen.DrawImage(boardImage, nil)

	emptyImage, _ := ebiten.NewImage(tileSize, tileSize, ebiten.FilterDefault)
	purple := color.RGBA{0xff, 0x0, 0xff, 0xff}
	emptyImage.Fill(purple)
	offset := float64(tileSize) * float64(0.05)
	for i := 0; i < 4; i++{
		for j:= 0; j < 4; j++{
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(0.9, 0.9)
			op.GeoM.Translate(offset+float64(tileSize*i), offset+float64(tileSize*j))
			screen.DrawImage(emptyImage, op)
		}
	}
	// ebitenutil.DebugPrint(screen, "Hello, World!")
	return nil
}

func main() {
	board := NewBoard(0, 0, 10, 4)
	fmt.Println(board)
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}
