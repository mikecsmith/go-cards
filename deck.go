package main

import "fmt"

// Create a new type of deck
// which is a slice of strings

type deck []string

// Creates a new deck
func newDeck() deck {
	cardSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	var cards deck
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
