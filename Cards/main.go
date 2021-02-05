package main

import "fmt"

func main() {
	// INITIALIZE NEW DECK
	cards := newDeck()
	//SHUFFLE DECK
	cards.shuffle()
	// fmt.Println("len: ", len(cards))
	//DEAL a hand from deck
	var hand deck
	hand, cards = deal(cards, 5)
	fmt.Println("####HAND####")
	hand.print()
	fmt.Println("hand len: ", len(hand))
	// fmt.Println("####DECK####")
	// cards.print()
	// fmt.Println("deck len: ", len(cards))
	//save hand to file
	hand.saveToFile("test.txt")
	//load from file
	loaded := loadFromFile("test.txt")
	fmt.Println("Hand loaded from file")
	loaded.print()
}
