//testing slice parameters
package main

import "fmt"

func main() {
	slice := []string{"a", "b", "c"}
	fmt.Println(slice)
	//note passing slice to Test allow slice modification
	Test(slice)
	fmt.Println(slice)
}

func Test(s []string) {
	s[0] = "zoot"
}
