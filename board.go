package main

import (
	"fmt"
	"image/color"
	"github.com/hajimehoshi/ebiten"
)

var (
	tileColor = color.RGBA{0xff, 0xff, 0xf0, 0xff}
)

type gameObject struct {
	posX, posY int
}

func (g gameObject) String() string {
	return fmt.Sprintf("(%v, %v)", g.posX, g.posY)
}

// Piece is a player's piece
type Piece struct {
	gameObject
}

// Arrow object
type Arrow struct {
	gameObject
}

// Board of a game
type Board struct {
	gameObject
	Pieces     [][]*Piece
	tileSize   int
	tileNum    int
}

// NewBoard creates a new gameboard
func NewBoard(posx, posy, tilenum, tilesize int) Board {
	return Board{
		gameObject: gameObject {posX: posx, posY: posy},
		Pieces: 	make([][]*Piece, tilenum),
		tileNum:  	tilenum,
		tileSize: 	tilesize,
	}
}

// Initialize state of a board
func (b Board) Initialize(mask [][]int) {
	// b.Pieces := make([][]Piece, b.tileNum)
	for i := 0; i < b.tileNum; i++{
		b.Pieces[i] = make([]*Piece, b.tileNum)
		for j := 0; j < b.tileNum; j++{
			switch mask[i][j] {
			case 0:
				b.Pieces[i][j] = nil
			case 1:
				b.Pieces[i][j] = &Piece{gameObject{i, j}}
			}
			}
	}
}

func (b Board) String() string {
	return fmt.Sprintf("x: %v, y: %v, \npieces: %v", b.posX, b.posY, b.Pieces)
}

// SelectTile of a gameboard
func (b Board) SelectTile(x, y int) {
	i := (x - b.posX) / b.tileSize
	j := (y - b.posY) / b.tileSize
	fmt.Printf("i: %v, j: %v\n", i, j)
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
