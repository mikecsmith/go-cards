package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"regexp"
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

// Utility function to join a deck into a string and then typecast to a byte slice
func deckToString(d deck) []byte {
	return []byte(strings.Join(d, ","))
}

// Utility function to convert a deck in a byte slice JSON
func deckToJSON(d deck) []byte {
	// Marshal is the function to convert a data structure into a JSON byte slice (ASCII conversion)
	// It returns a byte slice json string and an error
	b, err := json.Marshal(d)

	// Utility function defined in utils.go - checks for an error in JSON conversion and panics if one is present
	check(err)

	return b
}

// Function to read the deck from a file
func readDeckFromFile(filename string) deck {
	// ReadFile ioutil method - returns a byte slice and an error
	b, err := ioutil.ReadFile(filename)
	// Check the error - panic if error not nil
	check(err)

	// Compile a regexp object to check for files of type .json
	r := regexp.MustCompile(".*?\\.json")
	// Check to see if the filename is of type json
	if r.MatchString(filename) {
		// Declare an empty variable of type deck
		var d deck
		// Unmarshal the json - unmarshal takes a byte slice and a pointer to a variable then typecasts the json into that type - or errors if the type is not suitable
		err := json.Unmarshal(b, &d)
		// Check error and panic if not nil
		check(err)
		// return the json decoded deck
		return d
	}
	// If deck is not saved as a JSON assume deck is saved as comma delimited string, split the string into a []string then typecast to a deck
	return deck(strings.Split(string(b), ","))
}

// Save a deck to a JSON file
func saveDeckToFile(filename string, b []byte) {
	// Writes the deck to the specified filename in the current dir
	err := ioutil.WriteFile("./"+filename, b, 0644)

	// Check to see if error returned from WriteFile and panic if error present
	check(err)
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
