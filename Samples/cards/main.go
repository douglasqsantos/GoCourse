package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades" // := used to initialize the variable.
	// card = "Five of Diamonds" // When we already defined the variable to assigned another value just use =
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
