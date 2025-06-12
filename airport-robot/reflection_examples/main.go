package main

import (
	"fmt"
	"reflect"
)

type Example interface {
	Method1(int)
}

type Type1 struct{}

func (t Type1) Method1(int) {}

type Type2 struct{}

func (t Type2) Method1(int) {}

var types []reflect.Type

func getTypes() []reflect.Type {
	var types []reflect.Type
	interfaceType := reflect.TypeOf((*Example)(nil)).Elem()

	for i := 0; i < interfaceType.NumMethod(); i++ {
		method := interfaceType.Method(i)
		if method.Type.NumIn() == 0 {
			continue
		}
		receiverType := method.Type.In(0)
		if receiverType.Kind() == reflect.Ptr {
			receiverType = receiverType.Elem()
		}
		if receiverType.Kind() == reflect.Struct && reflect.PtrTo(receiverType).Implements(interfaceType) {
			types = append(types, receiverType)
		}
	}

	return types
}

func init() {
	types = getTypes()
}
func make(e Example) interface{} {
	fmt.Println(reflect.TypeOf(e))
	if f, ok := factory[fmt.Sprintf("%T", e)]; ok {
		ex := f()
		fmt.Printf("%T", ex)
		return ex
	}
	return nil
}

var factory map[string]func() interface{} = map[string]func() interface{}{
	"Type1": func() interface{} { return &Type1{} },
	"Type2": func() interface{} { return &Type2{} },
}

func New(g Example) interface{} {
	fmt.Println(reflect.TypeOf(g))
	if f, ok := factory[fmt.Sprintf("%T", g)]; ok {
		ex := f()
		fmt.Printf("%T", ex)
		return ex.(Example)
	}
	fmt.Println("Zoot")
	return nil
}

func type_from_reflect(e Example) {
	t := reflect.TypeOf(e)
	r := reflect.New(t).Interface()
	fmt.Printf("%T", r)
}

func main() {
	type_from_reflect(Type1{})
	fmt.Println(types)
	t1 := make(Type1{})
	t2 := make(Type2{})
	fmt.Printf("%T %T\n", t1, t2)
	t3 := New(Type1{})
	t4 := New(Type2{})
	fmt.Printf("%T %T\n", t3, t4)
}
