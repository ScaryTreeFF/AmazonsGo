package main

import (
	"fmt"
	"image/color"
	"github.com/hajimehoshi/ebiten"
)

var (
	tileColor = color.RGBA{0xff, 0xff, 0xf0, 0xff}
)

// Board of a game
type Board struct {
	PosX, PosY float64
	Pieces     []Piece
	tileSize   int
	tileNum    int
}

// NewBoard creates a new gameboard
func NewBoard(posx, posy float64, tilenum, tilesize, n int) Board {
	return Board{
		PosX:   posx,
		PosY:   posy,
		Pieces: make([]Piece, n),
		tileNum:  tilenum,
		tileSize: tilesize,
	}
}

func (b Board) String() string {
	return fmt.Sprintf("x: %v, y: %v, \npieces: %+v", b.PosX, b.PosY, b.Pieces)
}

// Draw a gameboard
func (b Board) Draw(screen *ebiten.Image) {
	boardImage, _ := ebiten.NewImage(screenWidth, screenHeight, ebiten.FilterDefault)
	boardImage.Fill(color.NRGBA{0x00, 0x40, 0x80, 0xff})
	screen.DrawImage(boardImage, nil)

	tileImage, _ := ebiten.NewImage(b.tileSize, b.tileSize, ebiten.FilterDefault)
	tileImage.Fill(tileColor)
	offset := float64(screenWidth - b.tileSize * b.tileNum) / 2
	for i := 0; i < b.tileNum; i++{
		for j:= 0; j < b.tileNum; j++{
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(0.98, 0.98)
			op.GeoM.Translate(offset+float64(b.tileSize*i), offset+float64(b.tileSize*j))
			screen.DrawImage(tileImage, op)
		}
	}
}
