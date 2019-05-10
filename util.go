package main

import (
	"strings"
	"fmt"
)

// Matrix2String converts matrix of pieces to a string representation
func Matrix2String(t [][]*Piece) string {
    var s []string
    for _, i := range t {
        s = append(s, fmt.Sprintf("%v", i))
    }
    return strings.Join(s, "\n")
}