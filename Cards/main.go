package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	fmt.Println("len: ", len(cards))
	var hand deck
	hand, cards = deal(cards, 5)
	fmt.Println("####HAND####")
	hand.print()
	fmt.Println("hand len: ", len(hand))
	fmt.Println("####DECK####")
	cards.print()
	fmt.Println("deck len: ", len(cards))
}
func newCard() string {
	return "5 of diamonds"
}
