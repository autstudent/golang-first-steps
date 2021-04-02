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

## Byte slices

A byte slice is a set of bytes which represents a certain amount of data in a specific format.

[10 234 123 45 224 111 231 123] -> Each number represents a character in ascii representation

This type or structure of data is so popular in golang to work with data or strings.

- Translate string to byte slide

```$bash
[]byte("Hi there!")
```

- Translate a slice to byte slide

```$bash
array := ["pepe","manuela"]
chain := strings.Join([]string(array), ","]
[]byte(chain)
```

## Handle errors

Normally, errors are a specific returned value in different functions. You need to review this value in order to detect errors:

```$bash
// bs -> new []byte from the file with the information
// err -> []byte with the error message
bs, err := ioutil.ReadFile(filename)
if err != nill {
    fmt.Println("Error:", err)
    os.Exit(1)
}
```

IMPORTANT: When you receive an error, it is possible to return a default value o quit the program. It is important to keep in mind this aspect when you implement a new func.

## Write and Read to/from file

- Write

```$bash
// File Name, a string in byte[] and permission for this file
return ioutil.WriteFile(filename, []byte(d.toString()), 0640)
```

- Read

```$bash
// bs -> new []byte from the file with the information
// err -> []byte with the error message
bs, err := ioutil.ReadFile(filename)
...
return string(bs)
```

## Change number position in slide

With this method, it is possible to change the specific item in a slice in a simple way:

- A non definitive algorithm:

```$bash
for i := range d {
    newPosition := rand.Intn(len(d)-1)
    d[i], d[newPosition] = d[newPosition], d[i]
}
```

- Definitive algorithm:

```$bash
source := rand.NewSource(time.Now().UnixNano())
r := rand.New(source)

for i := range d {
    newPosition := r.Intn(len(d) - 1)
    d[i], d[newPosition] = d[newPosition], d[i]
}
```

## Tests

Golang does not use a specific testing framework, usually this uses a normal golang code to test its applications.

To create a test, it is only required to create a new test file ending in **_test.gp** and run them with the following command:

```$bash
go test
```

Regarding implement a new test, it is required:

- Create a new file  **_test.gp**
- Added the regular package name (E.g. _package main_)
- Create a function per each function which we want to test wit UpperCase letters
- Added the function argument behind golang testing (_func TestNewDeck(t *testing.T)_)
- Added the expression adding if statement and triggering an error with *testing.T* library

```$bash
func TestNewDeck(t *testing.T) {
    d := newDeck()
    if len(d) != 16 {
        t.Error("Expected deck length of 16, but got", len(d))
    }
}
```
