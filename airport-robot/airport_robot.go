package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName() string
	Greet() string
}
type GermanGreeter struct{}

func (g GermanGreeter) LanguageName() string {
	return "German"
}
func (g GermanGreeter) Greet() string {
	return "Hallo"
}

type Italian struct{}

func (g Italian) LanguageName() string {
	return "Italian"
}
func (g Italian) Greet() string {
	return "Ciao"
}

type Portuguese struct{}

func (g Portuguese) LanguageName() string {
	return "Portuguese"
}
func (g Portuguese) Greet() string {
	return "Olá"
}

func SayHello(name string, g Greeter) string {
	return "I can speak " + g.LanguageName() + ": " + g.Greet() + " " + name + "!"
}
