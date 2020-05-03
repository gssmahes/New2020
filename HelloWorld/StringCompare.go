package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("India", "Indiana"))
	fmt.Println(strings.Compare("Indianjjjjjjjja", "India"))
	fmt.Println(strings.Compare("India", "India"))

	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))
}
