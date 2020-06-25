package main

func main() {
	cards := newDeck()
	shuffledCards := cards.shuffle()

	hand, deck := shuffledCards.deal(5)

	hand.print()
	deck.print()

}
