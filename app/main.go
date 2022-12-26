package main

import (
	"fmt"
	"net/http"
	"time"
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

	for l := range ch {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, ch)
		}(l)
	}
}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link + " is down!")
		ch <- link
		return
	}

	fmt.Println(link + " is up and running!")
	ch <- link
}
