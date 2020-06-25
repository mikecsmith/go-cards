package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Create a new type of deck
// which is a slice of strings

type deck []string

// Creates a new deck
func newDeck() deck {
	var cards deck

	cardSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Prints all the cards in a deck
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// Deals n cards
func (d deck) deal(handSize int) deck {
	// Seed the rand function so it does not always return the same values
	rand.Seed(time.Now().UnixNano())

	// Define the startNumber for the deck by generating a random number in the range 0 to the length of the deck, minus the length of the handSize
	startNumber := rand.Intn(len(d) - handSize)

	// Return a range from the deck, starting at the startNumber and ending at the startNumber + the handSize
	return d[startNumber : startNumber+handSize]
}

// Shuffle the deck
func (d deck) shuffle() deck {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// rand.Shuffle accepts an integer which represents the number of elements to shuffle and a swap function that swaps the values of i & j - note: the swap function mutates d
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})

	// return the shuffled deck
	return d
}
