package main

import (
	"fmt"
)

func main() {
	var number, remainder, temp int
	var reverse int = 0

	fmt.Println("Enetr any number")
	fmt.Scan(&number)

	fmt.Println("Entered number is ", number)
	temp = number

	for {
		remainder = number % 10
		reverse = reverse*10 + remainder
		number = number / 10

		if number == 0 {
			break
		}
	}
	if temp == reverse {
		fmt.Println("Entered number is palindrome")
	} else {
		fmt.Println("Entered number is NOT palindrome")
	}

}
