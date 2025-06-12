package main

import (
	hamming "exercism/Go/hamming/code"
	"fmt"
)

func main() {

	strand1 := "GGACGGATTCTG"
	strand2 := "GGCCGGATCCTA"

	fmt.Println(strand1)
	fmt.Println(strand2)
	dist, _ := hamming.Distance(strand1, strand2)
	fmt.Printf("Hamming distance: %v", dist)
}