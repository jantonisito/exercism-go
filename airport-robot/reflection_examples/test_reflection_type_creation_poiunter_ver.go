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

func make_inst_from_type_by_reflect_and_use_it(e Example) {
	t := reflect.TypeOf(e)
	r := reflect.New(t).Interface()

	// Check if the type stored in r implements the Example interface
	if reflect.TypeOf(r).Elem().Implements(reflect.TypeOf((*Example)(nil)).Elem()) {
		ptr := reflect.ValueOf(r).Elem()
		useIt(ptr.Interface().(*Example))
	}
}

func useIt(e *Example) {
	fmt.Printf("%T\n", e)
}

func main() {
	make_inst_from_type_by_reflect_and_use_it(Type1{})
	// t := new(Type1)
	// fmt.Printf("%T\n", t)
	// useIt(t)
}
