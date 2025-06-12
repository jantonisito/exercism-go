package main

import "fmt"
import "time"

func main() {
	layout := "2006/01/02"
	date := "2008/11/11"
	t, _ := time.Parse(layout, date)
	fmt.Println(t)
}
