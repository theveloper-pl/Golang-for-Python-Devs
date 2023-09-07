package main

import (
	"fmt"
	"time"
)

func main() {

	// Niebuforowany kanał (Unbuffered channel) nie jest tym samym co buforowany kanał (Buffered channel) z limitem 1. Mimo że oba mogą "przechowywać" tylko jeden element, działają inaczej.

	// my_channel := make(chan int) // Ten kanał spowoduje błąd, ponieważ nie ma gorutyny, która z niego odbierałaby dane.
	my_channel := make(chan int, 1) // Buforowany kanał nie spowoduje błędu, ponieważ wysłana wiadomość będzie czekać w kolejce.

	fmt.Println("Wysyłanie do kanału my_channel")
	my_channel <- 1
	fmt.Println("Wysłano do kanału my_channel")


	time.Sleep(time.Second * 2)


	fmt.Println("Odbieranie z kanału my_channel")
	<-my_channel
	fmt.Println("Odebrano z kanału my_channel")

}
