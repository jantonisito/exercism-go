// outdated code - check: working_reflection_ex.go
package main

import (
	"fmt"
	"reflect"
)

type Example interface {
	Method1()
}

type Type1 struct{}

func (t Type1) Method1() {}

type Type2 struct{}

func (t Type2) Method1() {}

func makeIt(e Example) {
	t := reflect.TypeOf(e)
	r := reflect.New(t).Interface()
	fmt.Printf("in_make_.. :%T\n", r)
	rv := reflect.ValueOf(r).Elem()
	//x := rv.Interface().(Example)
	//useIt(x)
	useIt(rv.Addr().Interface().(*Example))
}

//	func useIt(e Example) {
//		fmt.Printf("in useIt: %T\n", e)
//	}
func useIt(e *Example) {
	fmt.Printf("in useIt: %T\n", e)
}

func main() {
	makeIt(Type1{})
}
