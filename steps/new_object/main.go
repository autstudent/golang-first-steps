package main

func main() {
	// Create a Deck
	cards := newDeck()

	// Print cards
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
