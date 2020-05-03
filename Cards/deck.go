package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type ceetu []string

func newDeck() ceetu {
	cards := ceetu{}
	cardSuits := []string{"spades", "diamonds", "hearts", "clubs"}
	cardValues := []string{"1", "2", "3", "4"}

	for _, suits := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, values+" of "+suits)
		}
	}
	return cards
}

func (d ceetu) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d ceetu, handSize int) (ceetu, ceetu) {
	a, b := d[:handSize], d[handSize:]
	return a, b
}

func (d ceetu) toString() string {
	return strings.Join(d, ",")
}

func (d ceetu) saveToDrive(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)

}

func newDeckFromDrive(fileName string) (d ceetu) {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error1:", err)
		os.Exit(1)

	}
	s := strings.Split(string(bs), ",")
	return s

}
