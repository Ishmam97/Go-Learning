package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
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
	if handSize >= 0 {
		return d[:handSize], d[handSize:]
	} else if handSize > len(d) {
		fmt.Print("cannot deal more cards than deck has")
		return d, nil
	} else {
		fmt.Print("Error occured")
		return d, nil
	}

}
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
func (d deck) toString() string {
	//type casting
	return strings.Join([]string(d), ",")
}
func (d deck) toByte() []byte {
	//type casting
	return []byte(d.toString())
}
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, d.toByte(), 0666)
}
func loadFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error occured: ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}
