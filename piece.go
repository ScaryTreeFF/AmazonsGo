package main

import (
	"fmt"
)

// Piece is a game piece
type Piece struct {
	PosX, PosY float64
}

func (p Piece) String() string {
	return fmt.Sprintf("(%v, %v)", p.PosX, p.PosY)
}