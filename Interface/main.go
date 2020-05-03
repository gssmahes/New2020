package main

import "fmt"

type englishBot struct{}
type tamilBot struct{}

type bot interface {
	getGreetings() string
}

func main() {
	eb := englishBot{}
	tb := tamilBot{}

	printGreetings(eb)
	printGreetings(tb)
}

func (englishBot) getGreetings() string {
	return "Hi"
}

func (tamilBot) getGreetings() string {
	return "Vanakkam"
}

func printGreetings(b bot) {
	fmt.Println(b.getGreetings())
}
