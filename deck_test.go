package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	l := len(d)

	if l != 52 {
		t.Errorf("Expected deck length of 52, but got %v", l)
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected first card to equal Ace of Clubs, but got %v", d[0])
	}

	if d[l-1] != "King of Spades" {
		t.Errorf("Expected last card to equal King of Spades, but got %v", d[l-1])
	}
}

//Tests saving and loading - important to do cleanup yourself with Go as it will not hold your hand
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 decks in card, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
