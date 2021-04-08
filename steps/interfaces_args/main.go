package main

import (
	"fmt"
	"strconv"
)

type bot interface{
	getGreeting(int, string, int) (string)
}

type englishBot struct {}
type spanishBot struct {}

func main()  {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting(a int, b string, c int) string  {
	// BUSINESS LOGIC Here...
	return "Hi There!" + strconv.Itoa(a) + b + strconv.Itoa(c)
}

func (spanishBot) getGreeting(a int, b string, c int) string  {
	// BUSINESS LOGIC Here...
	return "Hola!" + strconv.Itoa(a) + b + strconv.Itoa(c)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting(12,"eg",2))
}