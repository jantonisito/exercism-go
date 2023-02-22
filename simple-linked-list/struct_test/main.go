package main

import (
	"fmt"
)

type Foo struct {
	Bar string
}

func (f *Foo) String() string {
	return f.Bar
}

func (f *Foo) Set(s string) {
	f.Bar = s
}

func (f *Foo) Default() {
	f.Bar = "default"
}

func main() {
	f := &Foo{Bar: "foo"}
	fmt.Println(f.String())
	f.Set("Hello")
	fmt.Println(f.String())
	g := &Foo{}
	g.Default()
	fmt.Println(g.String())

}
