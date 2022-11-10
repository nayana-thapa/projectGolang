package main

func main() {

	// cards := newDeck()
	// fmt.Println(cards.saveToFile("my_cards"))
	cards := newDeck()
	cards.shuffle()
	cards.print()

}
