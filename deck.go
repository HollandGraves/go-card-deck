package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck
// which is a slice of strings

type deck []string

// Create a function that creates
// a deck

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Clubs", "Spades", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Create a function that
// prints a deck type out

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Create a function that
// creates a hand of cards

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Create a function that
// turns the items of a
// deck type into a list
// of strings

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Create a function that
// saves a string to file

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Create a function that
// reads a file on the hard
// drive

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Common options to handle an err in this situation
		// Option #1 - Log the error and return a call to newDeck()
		// Option #2 - Log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

// Create a function that
// generates a random number
// and loops through each index
// and swaps that index with
// random number

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	r.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}
