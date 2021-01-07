package main

import "fmt"

type bot interface {
	getGreeting() string
}

type engBot struct{}
type spaBot struct{}

func main() {
	eb := engBot{}
	sb := spaBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (engBot) getGreeting() string {
	return "Hi!"
}

func (spaBot) getGreeting() string {
	return "Hola!"
}
