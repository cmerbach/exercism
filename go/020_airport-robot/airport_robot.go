package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(string) string
}

func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}


type Italian struct {}


func (g Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

type Portuguese struct {}


func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (i Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}