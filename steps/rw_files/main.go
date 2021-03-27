package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
	cardsNew := newDeckFromFile("my_cards")
	cardsNew.print()
}