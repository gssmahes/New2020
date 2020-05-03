package main

import "fmt"

func main() {

	//map declaration1 - map[key]value
	var game map[string]string
	fmt.Print("map declaration1 is")
	fmt.Println(game)

	//map declaration2
	game2 := make(map[string]string)
	fmt.Println("map declaration2 is")
	game2["cric"] = "Cricket"
	fmt.Println(game2)

	//map deletion
	fmt.Println("map declaration2 - cricket deletion")
	delete(game2, "cric")
	fmt.Println(game2)

	//map declaration3
	fmt.Println("Map declaration3")
	cricket := map[string]string{
		"batsman":    "sachin",
		"bowler":     "kumble",
		"allrounder": "Yuvraj",
	}

	cricket["fielder"] = "Robin singh"

	mapPrint(cricket)
}

func mapPrint(c map[string]string) {
	for key, value := range c {
		fmt.Println("Best Indian ", key, "is ", value)
	}
}
