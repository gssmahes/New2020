package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.in",
		"http://gmail.com",
		"http://twittekurtr.com",
		"http://golang.org",
		"http://amazonrt.in",
	}

	for _, link := range links {
		checkLinks(link)
	}
}

func checkLinks(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is down")
		return
	}

	fmt.Println(url, "is up")
}
