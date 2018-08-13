package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	first_card := Card{"Spades", "Ace"}
	last_card := Card{"Hearts", "Four"}

	//confirm length
	if len(d) != 16 {
		t.Errorf("Expected deck to have 16  cards, but got %v cards", len(d))
	}
	// //confirm 1st card Ace of Spades
	if d[0] != first_card {
		t.Errorf("Expected first card to be Ace of Spades, but got %s of %s", d[0].value, d[0].suit)
	}
	// //confirm 1st card Ace of Spades
	if d[len(d)-1] != last_card {
		t.Errorf("Expected first card to be Four of Hearts, but got %s of %s", d[len(d)-1].value, d[len(d)-1].suit)
	}

}

func TestSaveLoadDeckFromFile(t *testing.T) {
	os.Remove("__decktest")
	d := newDeck()
	d.deckToFile("__decktest")
	new_d := deckFromFile("__decktest")
	os.Remove("__decktest")
	//fmt.Println(new_d[0] == d[0] && new_d[0] == d[0])
	if new_d[0] != d[0] || new_d[len(new_d)-1] != d[len(d)-1] || len(d) != len(new_d) {
		t.Errorf("Deck saved to file should be same as deck loaded from file")
	}

}
