package main

import "fmt"

func main() {
	card := "Ace of Spades" //Variable declaration - alternative is var card = "card" or var card string = "card"
	card = newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
