package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	deckSize := 16
	firstCard := "Ace of Spades"
	lastCard := "Four of Clubs"

	if len(d) != deckSize {
		t.Errorf("Expected deck length of %v, but got %d", deckSize, len(d))
	}

	if d[0] != firstCard {
		t.Errorf("Expected %v, but got %v", firstCard, d[0])
	}

	if d[len(d)-1] != lastCard {
		t.Errorf("Expected %v, but got %v", lastCard, d[len(d)-1])
	}
}
