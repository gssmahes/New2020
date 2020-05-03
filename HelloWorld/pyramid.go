package main

import "fmt"

func main() {
	var rows int
	fmt.Println("Enter the no of rows ")
	fmt.Scan(&rows)

	for i := 1; i <= rows; i++ {
		//leaving space
		for space := 1; space <= rows-i; space++ {
			fmt.Print("  ")
		}

		//enter row values like 1, 3, 5, 7
		k := 2*i - 1
		for j := 1; j <= k; j++ {
			fmt.Print("* ")
		}
		fmt.Println("")

	}
}
