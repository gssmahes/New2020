package main

import (
	"fmt"
	"net/http"
	"os"
	
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error is ", err)
		os.Exit(1)
	}

	//allocate no of elements to  byteslice
	byteslice := make([]byte, 9999)

	//Read interface push the readed info to byetslice. Resp means response from http.Get()

	resp.Body.Read(byteslice)

	fmt.Println(string(byteslice))

}
