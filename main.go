package main

import "fmt"

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	for index, item := range cards {
		fmt.Println(index, item)
	}

	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
