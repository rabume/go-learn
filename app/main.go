package main

import (
	"fmt"
)

func main() {
	// Create a map with the default "zero" values
	// Prints out: map[]
	// var colors1 map[string]string
	colors1 := make(map[int]string)

	// Add values later on to the above created maps
	colors1[1] = "#ffffff"
	colors1[2] = "#ff0000"

	delete(colors1, 2)

	// Create map and add directly some values
	colors := map[string]string{
		"red":    "#ff0000",
		"yellow": "#ffdd00",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
