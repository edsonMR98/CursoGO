package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new Type of "deck", wich
// extends functionality of slice of string
type deck []string

// newDeck Create 16 combinations of this deck
// return a deck
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

// Receiver function, Print the whole deck
// d: deck to print
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// deal create two "deck", a hand and the remaining cards
// d: a deck to deal
// handsize: size of the new hand
// return a hand and remaining cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Receiver function, toString become a deck to string
// return a deck like a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Receiver function, saveToFile save a deck in our system files
// fileString: name of the new created file
// return a error, if it´s the case
func (d deck) saveToFile(fileString string) error {
	return ioutil.WriteFile(fileString, []byte(d.toString()), 0666)
}
