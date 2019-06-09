package main

import "fmt"

func main() {
	cards := newDeck()

	cards.shuffle()
	fmt.Println("########Repartir")
	fmt.Println("################")

	hand, _ := deal(cards, 5)
	hand.Print()

	fmt.Println(hand.toString())

	fmt.Println("########Deck From File")

}
