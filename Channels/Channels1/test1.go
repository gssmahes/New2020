package main

import (
	"fmt"
)

func printDay(boolChan chan bool){
	fmt.Println("Saturday")
	boolChan <- true

}

func main(){
	boolChan := make(chan bool)
	go printDay(boolChan)
	<-boolChan
	fmt.Println("end of main function")

}