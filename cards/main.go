package main

func main() {
	// Create a Deck
	cards := newDeck()

	// Split the deck
	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()

}