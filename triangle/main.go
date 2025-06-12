package main

import (
	triangle "exercism/Go/triangle/code"
	"fmt"
)

func main() {
	fmt.Println(triangle.KindFromSides(1, 1, 1))
	fmt.Println(triangle.KindFromSides(1, 1, 3))
	fmt.Println(triangle.KindFromSides(2, 1, 2))
}
