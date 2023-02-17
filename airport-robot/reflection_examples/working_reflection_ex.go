package main

import (
	"fmt"
	"reflect"
)

type Example interface {
	Method1()
}

type Type1 struct {
	foo int
}

func (t Type1) Method1() {}

type Type2 struct {
	foo int
}

func (t Type2) Method1() {}

func makeIt(e Example) {
	t := reflect.TypeOf(e)
	r := reflect.New(t).Interface()
	rv := reflect.ValueOf(r).Elem()
	x := rv.Addr().Interface().(Example)

	fmt.Printf("type in_make_.. :%T\n", x)
	fmt.Println("x.foo before call to useIt: ", x.(*Type1).foo)
	useIt(x)
	fmt.Println("x.foo after call to useIt: ", x.(*Type1).foo)
	if ok := setIt(x); ok != nil {
		fmt.Println(ok.Error())
	}
	fmt.Println("x.foo after call to setIt: ", x.(*Type1).foo)
}

func useIt(e Example) {
	fmt.Printf("type in useIt: %T\n", e)
	x := e.(*Type1)
	fmt.Println("x.foo in useIt: ", x.foo)
	x.foo = 2
	fmt.Println("x.foo in useIt: ", x.foo)
}

func setIt(e Example) error {
	gValue := reflect.ValueOf(e)
	lField := gValue.FieldByName("foo")
	if !lField.IsValid() {
		return fmt.Errorf("type %T does not have a field 'foo'", e)
	}
	if !lField.CanSet() {
		return fmt.Errorf("field 'foo' of type %T is not settable", e)
	}
	lField.SetInt(17)
	return nil
}

func main() {
	makeIt(Type1{})
}
