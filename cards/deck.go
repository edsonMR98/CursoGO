package main

import "fmt"

// Create a new Type of "deck", wich
// extends functionality of slice of string
type deck []string

func newDeck() deck {
	cards := deck{}
	cardsSuits := []string{"Diamonds", "Spades", "Hearts", "Clubs"}
	cardsValues := []string{"One", "Two", "Three", "Four"}

	for _, suit := range cardsSuits {
		for _, value := range cardsValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
