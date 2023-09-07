package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// https://levelup.gitconnected.com/how-to-implement-concurrency-and-parallelism-in-go-83c9c453dd2


func checkWebsiteStatus(url string, wg *sync.WaitGroup) {
	defer wg.Done()  // Oznacza zakończenie tej goroutine po jej wykonaniu

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

	var wg sync.WaitGroup
	wg.Add(len(websites))  // Dodaj liczbę stron do WaitGroup

	// Z Goroutines
	for _, website := range websites {
		go checkWebsiteStatus(website, &wg)
	}

	wg.Wait()  // Czekaj na zakończenie wszystkich goroutines

	elapsedTime := time.Since(startTime)
	fmt.Printf("Sprawdzanie zakończone w %s\n", elapsedTime)
}
