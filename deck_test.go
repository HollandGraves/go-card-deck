package main

import (
	"testing"
)

// these following shorthand codes are placeholders for the following specific value
// e.g.
//
// %d is the placeholder for an integer
// %s is the placeholder for a string
// %+v is the placeholder for every value an object has

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Error: deck of cards is not 52, it is %d", len(d))
	}
}
