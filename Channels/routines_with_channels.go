package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.in",
		"http://gmail.com",
		"http://twitter.com",
		"http://golang.org",
		"http://amazon.in",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLinks(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLinks(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is down")
		c <- url
		return
	}

	fmt.Println(url, "is up")
	c <- url
}
