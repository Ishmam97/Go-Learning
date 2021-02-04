package main

import (
	"fmt"
	"math/rand"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Clovers", "Diamonds"}
	cardVals := []string{"ACE", "KING", "QUEEN", "JACK", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	for _, suit := range cardSuits {
		for _, val := range cardVals {
			cards = append(cards, val+" of "+suit)
		}
	}
	return cards
}
func (d deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
}
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
