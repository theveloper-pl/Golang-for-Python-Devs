package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkWebsiteStatus(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(url + " is down!")
		return
	}

	defer resp.Body.Close()
	fmt.Println(url + " is up with status code: " + fmt.Sprint(resp.StatusCode))
}

func main() {
	startTime := time.Now()

	websites := []string{
		"https://www.google.com",
		"https://www.openai.com",
		"https://www.reddit.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://www.golang.org",
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.google.com",
		"https://www.openai.com",
		"https://www.reddit.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://www.golang.org",
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.google.com",
		"https://www.openai.com",
		"https://www.reddit.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://www.golang.org",
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.google.com",
		"https://www.openai.com",
		"https://www.reddit.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://www.golang.org",
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.google.com",
		"https://www.openai.com",
		"https://www.reddit.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://www.golang.org",
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.google.com",
		"https://www.openai.com",
		"https://www.reddit.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://www.golang.org",
		"https://www.youtube.com",
		"https://www.amazon.com",
	}


	for _, website := range websites {
		checkWebsiteStatus(website)
	}


	elapsedTime := time.Since(startTime)
	fmt.Printf("Sprawdzanie zakończone w %s\n", elapsedTime)
}
