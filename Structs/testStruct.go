package main

import "fmt"

type cricRest struct {
	allRounder string
	keeper     string
	captain    string
}

type cricketers struct {
	batsman string
	bowler  string
	CR      cricRest
}

func main() {
	indian := cricketers{
		batsman: "sachin",
		bowler:  "Zaheer",
		CR: cricRest{
			allRounder: "Yuvraj",
			keeper:     "Dinesh",
			captain:    "Dhoni",
		},
	}
	//indianBatsman := &indian
	(&indian).updateBatsman("Sachin")
	indian.Print()
}

func (c *cricketers) updateBatsman(bm string) {
	c.batsman = bm
}

func (c cricketers) Print() {
	fmt.Printf("%+v", c)
}
