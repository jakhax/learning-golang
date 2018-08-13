package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//card struct
type Card struct {
	suit, value string
}

//deck slice
type deck []Card

func newDeck() deck {
	//create a list of playing cards
	//essentially an array of strings
	cards := deck{}
	//we will use an unconvetional way to create a basic deck
	card_suits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	card_values := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range card_suits {
		for _, value := range card_values {
			cards = append(cards, Card{suit, value})
		}
	}
	return cards
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) print() {
	for i, c := range d {
		fmt.Printf("%v %s of %s\n", i, c.value, c.suit)
	}
}

func (d deck) toString() string {
	s := []string{}
	for _, c := range d {
		s = append(s, c.value+" of "+c.suit)

	}
	return strings.Join([]string(s), ",")
}

func (d deck) deckToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
