package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()
	fmt.Println(cards.toString())
	cards.deckToFile("file")
	deckFromFile("file")
}
