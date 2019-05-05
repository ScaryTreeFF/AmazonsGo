package main

import (
	"fmt"
)

// Board of a game
type Board struct {
	PosX, PosY float64
	Pieces     []Piece
	TileSize   float64
}

// NewBoard creates a new gameboard
func NewBoard(posx, posy, tilesize float64, n int) Board {
	return Board{
		PosX:   posx,
		PosY:   posy,
		Pieces: make([]Piece, n),
		TileSize: tilesize,
	}
}

func (b Board) String() string {
	return fmt.Sprintf("x: %v, y: %v, \npieces: %+v", b.PosX, b.PosY, b.Pieces)
}
