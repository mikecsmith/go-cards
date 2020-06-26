package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
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
func (d deck) deal(handSize int) (deck, deck) {
	// Return the hand and remaining deck
	return d[:handSize], d[handSize:]
}

// Save a deck to a txt file
func saveDeckToFileString(d deck) {
	// Take the deck and use the Join method from the strings standard library to join the deck into a single string using comma as delimiter
	stringifiedDeck := strings.Join(d, ",")
	// Typecast the stringifiedDeck into a byte slice
	byteDeck := []byte(stringifiedDeck)
	// Write the deck out to deck.txt
	err := ioutil.WriteFile("./deck.txt", byteDeck, 0644)
	// Check to see if error returned from WriteFile and panic if error present
	check(err)
}

// Save a deck to a JSON file
func saveDeckToFileJSON(d deck) {
	// Marshal is the function to convert a data structure into a JSON byte slice (ASCII conversion)
	// It returns a byte slice json string and an error
	b, err := json.Marshal(d)
	// Utility function defined in utils.go - checks for an error in JSON conversion and panics if one is present
	check(err)

	// Writes the deck to /tmp/deck
	err2 := ioutil.WriteFile("./deck.json", b, 0644)

	// Check to see if error returned from WriteFile and panic if error present
	check(err2)
}

// Shuffle the deck
func (d deck) shuffle() deck {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// rand.Shuffle accepts an integer which represents the number of elements to shuffle and a swap function that swaps the values of i & j
	// note: the swap function mutates d
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})

	// return the shuffled deck
	return d
}
