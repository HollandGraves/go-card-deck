package main

import (
	"os"
	"testing"
)

// these following shorthand codes are placeholders for the following specific value
// e.g.
//
// %d is the placeholder for an integer
// %s is the placeholder for a string
// %v is the placeholder for a single (the first) value that is returned by the function
// %+v is the placeholder for every value an object can return

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Error: deck of cards is not 52, it is %v", len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Error: first card of deck is not Ace of Clubs, it is %v", d[0])
	}

	if d[len(d)-1] != "King of Diamonds" {
		t.Errorf("Error: last card of deck is not King of Diamonds, it is %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Error: deck of cards is not 52, it is %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
