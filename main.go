package main

import "fmt"

type englishBot struct {
}

type spanishBot struct {
}

type frenchBot struct{}

type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	fb := frenchBot{}

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(fb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	//Very custom logic for generating an english gretting
	return "Hello, There!"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

func (frenchBot) getGreeting() string {
	return "Bonjour"
}
