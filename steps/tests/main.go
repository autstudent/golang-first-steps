package main

import (
	"fmt"
	"time"
)

func main() {
	cards := newDeck()
	cards.print()
	cards.shuffle()
	cards.print()

	fmt.Println(time.Now().UnixNano())
}