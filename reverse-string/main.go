package main

import (
	reverse "exercism/Go/reverse-string/code"
	"fmt"
)

func main() {
	var x string = "Hello, 世界 @&"
	fmt.Println(x)

	for i, ch := range reverse.Reverse(x) {
		fmt.Printf("%2d, %5T, %5v, %5c\n", i, ch, ch, ch) //left padding used
	}
	fmt.Println(reverse.Reverse(x))

	x = "reward"
	fmt.Println(reverse.Reverse(x))

}
