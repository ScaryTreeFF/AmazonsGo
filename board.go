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
	WhitePlayer bool
}

// Arrow object
type Arrow struct {
	gameObject
}

// Board of a game
type Board struct {
	gameObject
	Pieces     [][]*Piece
	Arrows	   []*Arrow // (???)
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
	// b.Pieces = make([][]*Piece, b.tileNum)
	for i := 0; i < b.tileNum; i++{
		b.Pieces[i] = make([]*Piece, b.tileNum)
		for j := 0; j < b.tileNum; j++{
			switch mask[i][j] {
			case 0:
				b.Pieces[i][j] = nil
			case 1:
				b.Pieces[i][j] = &Piece{gameObject{i, j}, true}
			case 2:
				b.Pieces[i][j] = &Piece{gameObject{i, j}, false}
			}
			}
	}
}

func (b Board) String() string {
	return fmt.Sprintf("x: %v, y: %v\npieces:\n%v", b.posX, b.posY, Matrix2String(b.Pieces))
}

// Coor2Ind return index of a tile
func (b Board) Coor2Ind(x, y int) (i, j int){
	i = (x - b.posX) / b.tileSize
	j = (y - b.posY) / b.tileSize
	fmt.Printf("i: %v, j: %v\n", i, j)
	return
}

// Ind2Coor return coord of a tile
func (b Board) Ind2Coor(i, j int) (x, y int){
	x = i * b.tileSize + b.posX
	y = j * b.tileSize + b.posY
	return
}


// Draw a gameboard
func (b Board) Draw(screen *ebiten.Image) {
	boardImage, _ := ebiten.NewImage(screenWidth, screenHeight, ebiten.FilterDefault)
	boardImage.Fill(color.NRGBA{0x00, 0x40, 0x80, 0xff})
	screen.DrawImage(boardImage, nil)

	tileImage, _ := ebiten.NewImage(b.tileSize, b.tileSize, ebiten.FilterDefault)
	tileImage.Fill(tileColor)
	
	for i := 0; i < b.tileNum; i++{
		for j:= 0; j < b.tileNum; j++{
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(0.98, 0.98)
			x, y := b.Ind2Coor(i, j)
			op.GeoM.Translate(float64(x), float64(y))
			screen.DrawImage(tileImage, op)
		}
	}
}
