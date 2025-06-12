package main

import (
	isogram "exercism/Go/isogram/code"
	"fmt"
)

func main() {
	str := "not isogram"
	var state = "is not"
	if isogram.IsIsogram(str) {
		state = "is"
	}
	fmt.Println(str, state, "isogram")

	str = "isogram"
	state = "is not"
	if isogram.IsIsogram(str) {
		state = "is"
	}
	fmt.Println(str, state, "isogram")

	digits := make([]byte, 0, 128)
	for i := 'A'; i <= 'z'; i++ {
		digits = append(digits, byte(i))
	}
	state = "is not"
	if isogram.IsIsogram(string(digits)) {
		state = "is"
	}
	fmt.Println(string(digits), state, "isogram")

}
