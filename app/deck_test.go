package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	deckSize := 16

	if len(d) != deckSize {
		t.Errorf("Expected deck length of %v, but got %d", deckSize, len(d))
	}
}
