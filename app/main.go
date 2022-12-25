package main

func main() {
	newCards := newDeck()
	newCards.saveToFile("my_cards.txt")

	cards := newDeckFromFile("my_cards.txt")
	cards.print()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
