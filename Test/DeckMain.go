package main

import "fmt"

func main() {
	cards := newCards()
	cards.Print()
	inLine, inHand := cards.deal(5)
	fmt.Println(".................In line......................")
	inLine.Print()
	fmt.Println(".................In hand......................")
	inHand.Print()
	cards.SaveToFile("MyDeck1")
	newCards := ReadFromFile("MyDeck1")
	fmt.Println("...................Reading from file....................")
	newCards.Print()
}
