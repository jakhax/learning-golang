package main

import (
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
