package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type germanBot struct{}

func main() {
	eb := englishBot{}
	sb := germanBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (eb englishBot) getGreeting() string {
	// Custom logic here
	return "Hello there!"
}

func (gb germanBot) getGreeting() string {
	// Custom logic here
	return "Guten Tag!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
