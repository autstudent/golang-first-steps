package main

import "fmt"

var car0 string

func main() {
	car0 = "Ace of spades - 0"

	var car1 string = "Ace of spades - 1"

	car2 := "Ace of spades - 2"
	car2 = "Ace of spades - 2 - Modified"

	car3 := newCard()

	fmt.Println(car0)
	fmt.Println(car1)
	fmt.Println(car2)
	fmt.Println(car3)
}

func newCard() string {
	return "Five of Diamonds"
}
