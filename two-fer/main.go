package main

import (
	//note that the import path is the same as the package name
	//twofer "exercism/Go/twofer/code"
	"fmt"
	//note that the import path is the same as the package name
	//but package does does not have to be called  "exercism/Go/twofer"
	twofer "zanzibar/code"
)

func main() {
	name := "Steve"
	fmt.Println(name)
	fmt.Println(twofer.ShareWith(name))
}
