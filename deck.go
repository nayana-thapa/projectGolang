package main

import "fmt"

// create a new type of 'deck'
// which is a slice if strings

type deck []string // this means deck == []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

// any variable of type 'deck' now gets access to the "print" method
// d is the actual of copy of the deck we're working with is availabe in the func
// as a variable called 'd'
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}

}
