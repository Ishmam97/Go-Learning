package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but got %v", len(d))
	}
	if d[0] != "ACE of Spades" {
		t.Errorf("Expected first card A of Spades got %v", d[0])
	}

}
func TestDeal(t *testing.T) {
	cards := newDeck()
	num := 5
	var hand deck
	hand, cards = deal(cards, num)
	if len(hand) != num {
		t.Errorf("length of hand error")
	}
	if len(cards) != 52-num {
		t.Errorf("length of deck error")
	}
}
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loaded := loadFromFile("_decktesting")
	if len(loaded) != 52 {
		t.Errorf("expected 52 cards got %v", len(loaded))
	}
	os.Remove("_decktesting")
}
