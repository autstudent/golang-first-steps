# Deeper into Go

We are working here with an application with different pieces:

- newDeck
- print
- shuffle
- deal
- saveToFile
- newDeckFromFile

## Variables

Basic go variables types in this course are float64, int, string and bool.

Define a new variable:

- var + var_name + var_type + = + value
- var_name + := + value

```$bash
...
// Declared variable
var card1 string = "Ace of spades - 1"
// Inferred variable
card2 := "Ace of spades - 2"

// Modify value
card2 = "Ace of spades - 2 - Modified"

// print
fmt.Println(card)
fmt.Println(card2)
...
```

## Functions

You need define the functions and declare the return type of value.

- func + func_name + () + return_type + {}

```$bash
func newCard() string {
    return "Five of Diamonds"
}
```

When you decided to support arguments, it is require to add them in the following way:

- func + func_name + ( arg_name arg_type, ) + return_type + {}

## Data structures

- Array -> Fixed length list of things
- Slice -> An array that can grow or shrink (*Must be of the same type of variable)

### Slice

```$bash
...
cards := []string
...
```

### For

Loops are very simple in Golang. You only need to declare a couple of variable behind for declaration and assign a range:

```$bash
...
cards := []string
for i, card := range cards {
    fmt.Println(i, card)
}
...
```

## Go Development approach

Go is not a Object oriented programming language. Basically, in Go the basics types (int, string, array, etc) are extended with a complex new "types" of data, for example _cards_. A function with a receiver is like a method, a function that belongs to an "instance".

- Define a new type and using it:

```$bash
cat deck.go
...
package main
// Create a new type "deck"
// which is a slice of strings
type deck []string
...
```

- Using it in a main go file:

```$bash
...
cards := deck{ "Ace of Diamonds", newCard()}
...
```

- Run the new golang app:

```$bash
go run ./main.go ./deck.go
```

## Receiver Functions

A receiver function is a function in GoLang with a receiver (any variable of type something which now gets access to a new method).

Basically, with receiver functions extend types with specific functions.

- Receiver function

```$bash
type deck []string

func (d deck) print() { 
    for i, card := range d {
        fmt.Println(i, card)
    }
}
```

- Receiver function call

```$bash
cards := deck{ "Ace of Diamonds", newCard()}
cards.print()
```

- Run the new golang app:

```$bash
go run ./main.go ./deck.go
```

NOTE: Normally, the variable name of the receiver is a letter or a couple of them (E.g "d" or "c")

## Return multiple values

It is possible to return multiple values in a function defining them in the function definition:

```$bash
func deal(d deck, handSize int) (deck, deck) {
    ...
    return var1, var2
}
```

- Receive both values

```$bash
hand, remainingDeck := deal(cards, 5)
```

## Tricks

- When you don't want to use a variable but you need to define it, you can replace it with "_"

```$bash
for i, value := range cardValues{
    cards = append(cards, value)
}
---
for _, value := range cardValues{
    cards = append(cards, suit+" of "+value)
}
```