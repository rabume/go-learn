package main

import (
	"fmt"
)

func main() {

	// Go will set the datatype int for this variable
	deckSize := 20
	fmt.Println(deckSize)

	// var card string = "Ace of spades"
	// Only need to use the ':' when we define a variable
	card := "Ace of Spades"

	// When assigning a new value to a already existing variable
	// you can just use '='
	card = "Five of Diamonds"

	fmt.Println(card)
}
