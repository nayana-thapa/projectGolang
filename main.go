package main

import "fmt"

func main() {

	cards := newDeck()
	fmt.Println(cards.saveToFile("my_cards"))

}
