package main

import (
	"fmt"
)

func main() {
	// my_channel := make(chan string)     // Unbuffered channel - Wiadomosc zostaje wyslana i musi byc odebrrana - Kod nie bedzie sie wykonywal dalej dopoki nie bedzie odbiorcy - Synchroniczna komunikacja
	my_channel := make(chan string, 3)  // Buffered channel - Queue - Wiadomosc zostaje wyslana, niema znaczenia to czy jest odbiorca czy nie - Asynchroniczna komunikacja

	chars := []string{"a", "b", "c"}

	for _, char := range chars {

		my_channel <- char

	}

	close(my_channel)

	for char := range my_channel {
		fmt.Println("Received:", char)
	}


	fmt.Println("Ending main")
}



