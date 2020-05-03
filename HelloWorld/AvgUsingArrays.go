package main

import "fmt"

func main() {
	var num [100]int
	var sum, n int
	var avg int = 0
	sum = 0
	fmt.Println("Enter  the number of elements")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Println("Enter the number")
		fmt.Scanln(&num[i])
		sum = sum + num[i]
	}

	avg = sum / n
	fmt.Println("the average is ", avg)

	fmt.Println("Length of an array ", len(num))
	fmt.Println("Capacity of an array", cap(num))



}
