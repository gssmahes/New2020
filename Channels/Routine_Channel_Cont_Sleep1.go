package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.in",
		"http://gmail.com",
		"http://twitter.com",
		"http://golangom.org",
		"http://amazon.in",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLinks(link, c)
	}

	for l := range c {
		//In child routine, dont pass actual variable such as l. so here, note (l) is passed to new variable link1 string here
		go func(link1 string) {
			time.Sleep(20*time.Second)
			checkLinks(link1, c)
		}(l)
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
