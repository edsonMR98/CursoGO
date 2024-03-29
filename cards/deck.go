package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
	"os"
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
// fileName: name of the new created file
// return a error, if it´s the case
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

// newDeckFromFile load a saved deck from system files.
// It becomes a bytes slice to string, to string slice and later to a deck
// fileName: name of the deck wich is saved in system files
// return the loaded deck
func newDeckFromFile(fileName string) deck {
	bytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	str := strings.Split(string(bytes), ",")

	return deck(str)
}

// Receiver function, shuffle a deck
// It uses random numbers to swap the position between two elements
// It creates a new seed to generate random numbers, using the current date
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	for x := range d {
		newPosition := random.Intn(len(d) - 1)
		d[x], d[newPosition] = d[newPosition], d[x]
	}
}