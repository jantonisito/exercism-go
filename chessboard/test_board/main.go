package main

import "fmt"

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard struct {
	board map[string]*File
}

func (b *Chessboard) Init() {
	b.board = make(map[string]*File)
	for _, row := range "ABCDEFGH" {
		pf := new(File)
		b.board[string(row)] = pf
	}
	//fmt.Println(*b)
}

func main() {
	board := Chessboard{}
	fmt.Println(board)
	board.Init()
	fmt.Println(board)
}
