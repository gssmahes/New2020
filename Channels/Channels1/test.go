package main

import (
	"fmt"
	"time"
)

func printDay(){
	fmt.Println("Saturday")
}

func main(){
	go printDay()
	time.Sleep(2*time.Second)
	fmt.Println("end of main function")

}