/*
	URL Checker used goroutines
*/
package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request fail")

type requestResult struct {
	url    string
	status string
}

func main() {
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	results := make(map[string]string)
	channel := make(chan requestResult)

	for _, url := range urls {
		go hitURL(url, channel)
	}

	for i := 0; i < len(urls); i++ {
		result := <-channel
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, channel chan requestResult) {
	response, err := http.Get(url)
	status := "OK"

	if err != nil || response.StatusCode >= 400 {
		status = "FAILED"
	}
	channel <- requestResult{url: url, status: status}
}
