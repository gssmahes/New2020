package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type ceetu []string

func newCards() ceetu {
	cards := ceetu{}
	cardSuits := []string{"clubs", "Diamond", "Heart", "Spades"}
	cardValues := []string{"ace", "2", "3", "4"}

	for _, suits := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, values+" of "+suits)
		}
	}
	return cards
}

func (allCards ceetu) Print() {
	for _, card := range allCards {
		fmt.Println(card)
	}
}

func (d ceetu) deal(handSize int) (ceetu, ceetu) {
	a, b := d[:handSize], d[handSize:]
	return a, b
}

func (c ceetu) SaveToFile(fileName string) {
	ioutil.WriteFile(fileName, []byte(strings.Join(c, ", ")), 0644)
}

func ReadFromFile(fileName string) ceetu {
	fs, _ := ioutil.ReadFile(fileName)
	return strings.Split(string(fs), ",")
}