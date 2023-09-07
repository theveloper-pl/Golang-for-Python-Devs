package main

import (
	"fmt"
	"time"
)

// do_something jest funkcją, która przyjmuje kanał tylko do odczytu.
// Wciąż działa, aż otrzyma wiadomość przez kanał is_done.
func do_something(is_done <-chan bool) {  // Kanał tylko do odczytu
	counter := 0
	for {
		select {
		case <-is_done:  // Oczekiwanie na wiadomość od kanału is_done
			fmt.Println("Done")
			return
		default:  // Działania domyślne, jeśli nie ma wiadomości w kanale
			fmt.Println("Doing something ", counter)
			time.Sleep(time.Second * 1)
			counter++
		}
	}
}

// do_something_different jest funkcją, która po pewnym czasie wysyła wiadomość do kanału is_done.
// W tym przypadku kanał jest tylko do zapisu.
func do_something_different(is_done chan<- bool) {  // Kanał tylko do zapisu

	time.Sleep(time.Second * 5)
	fmt.Println("Sending done message")
	is_done <- true  // Wysyłanie wiadomości do kanału

}

func main() {

	is_done := make(chan bool)  // Tworzenie kanału typu bool

	// Start dwóch gorutyn, które komunikują się za pomocą wspólnego kanału
	go do_something(is_done)  
	go do_something_different(is_done)

	time.Sleep(time.Second * 8)  // Główna gorutyna czeka przez 8 sekund, aby obie gorutyny mogły zakończyć swoje działania

}
