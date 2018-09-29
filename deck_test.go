package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		// %d is a
		// %s
		// %+v
		t.Errorf("Error: deck of cards is not 52, it is %d", len(d))
	}
}
