package airportrobot

import (
	"fmt"
	"reflect"
)

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName() string
	Greet() string
}
type GreeterData struct {
	language string
	greeting string
}

type greterDataMap map[Greeter]GreeterData

var greeter_data greterDataMap = greterDataMap{
	GermanGreeter{}: {"German", "Hallo"},
	Italian{}:       {"Italian", "Ciao"},
	Portuguese{}:    {"Portuguese", "Olá"},
}

func setGreeter(g *Greeter, data greterDataMap) error {
	gValue := reflect.ValueOf(*g)
	lField := gValue.FieldByName("Language")
	if !lField.IsValid() {
		return fmt.Errorf("type %T does not have a field 'Language'", *g)
	}
	if !lField.CanSet() {
		return fmt.Errorf("field 'Language' of type %T is not settable", *g)
	}
	lField.SetString(data[*g].language)

	gField := gValue.FieldByName("Greeting")
	if !gField.IsValid() {
		return fmt.Errorf("type %T does not have a field 'greeting'", *g)
	}
	if !gField.CanSet() {
		return fmt.Errorf("field 'language' of type %T is not settable", *g)
	}
	gField.SetString(data[*g].greeting)
	return nil
}

// func init() {
// 	for t, v := range greeter_data {
// 		fmt.Printf("%T\n", t)
// 		fmt.Println(v.language, v.greeting)
// 		err := setGreeter(t, v)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 		}
// 	}
//}

type GermanGreeter struct {
	Language string
	Greeting string
}

func (g GermanGreeter) LanguageName() string {
	return g.Language
}
func (g GermanGreeter) Greet() string {
	return g.Greeting
}

type Italian struct {
	Language string
	Greeting string
}

func (g Italian) LanguageName() string {
	return g.Language
}
func (g Italian) Greet() string {
	return g.Greeting
}

type Portuguese struct {
	Language string
	Greeting string
}

func (g Portuguese) LanguageName() string {
	return g.Language
}
func (g Portuguese) Greet() string {
	return g.Greeting
}

func SayHello(name string, g Greeter) string {
	t := reflect.TypeOf(g)
	r := reflect.New(t).Interface()
	rv := reflect.ValueOf(r).Elem()
	//rv := reflect.ValueOf(r)
	greeter := rv.Interface().(Greeter)

	setGreeter(&greeter, greeter_data)
	//setGreeter(rv.Interface().(Greeter), greeter_data)
	//setGreeter(rv.Interface().(*Greeter), greeter_data)
	return "I can speak " + g.LanguageName() + ": " + g.Greet() + " " + name + "!"
}
