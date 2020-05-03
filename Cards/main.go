package main

import "fmt"

func main() {
	cards := newDeck()
	//fmt.Println(cards)
	cards.print()
	fmt.Println("------------------new line-----------------------")
	hand, inCeetu := deal(cards, 5)
	hand.print()
	fmt.Println("--------------------In Hand------------------------")
	inCeetu.print()
	cards.saveToDrive("my_deck")
	fmt.Println("------------------Printing from Drive-----------------------")
	newCards := newDeckFromDrive("my_deck")
	fmt.Println(newCards)

}

//Functions to be written
//newDeck
//Print
//Deal
//saveToFile
//readFromFile
//toString
//Shuffle
//
