package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades" // := used to initialize the variable.
	// card = "Five of Diamonds" // When we already defined the variable to assigned another value just use =
	// card := newCard()
	// fmt.Println(card)

	// cards := []string{"Ace of Diamonds", newCard()}
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	// fmt.Println(cards)
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// fmt.Println("\t***** HAND *******")
	// hand.print()
	// fmt.Println("***** REMAINING *******")
	// remainingCards.print()
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
	// cards := newDeck()
	// cards.saveToFile("my_cards.txt")
	// fmt.Println(cards.toString())
	// cards := newDeckFromFile("my_cards.txt1")
	// cards.print()
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
