package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://golang.org",
		"https://rabu.me",
		"https://amazon.com",
		"https://google.ch",
	}

	ch := make(chan string)

	for _, link := range links {
		go checkLink(link, ch)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-ch)
	}
}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)

	if err != nil {
		ch <- link + " is down."
		return
	}

	ch <- link + " is up and running."
}
