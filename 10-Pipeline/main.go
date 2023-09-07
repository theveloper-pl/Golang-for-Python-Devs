package main

import (
	"fmt"
)

/*

W tym przypadku korzystamy z unbuffered channel. Co za tym idzie goroutina w funkcji generator wysyła jeden numer, 
a następnie blokuje się do momentu, aż goroutina w funkcji squared nie odbierze tego numeru. 
Sytuacja ta powtarza się dla każdego numeru aż do zamknięcia kanału.

*/

// generator: przekształca listę liczb całkowitych na kanał, który emituje te liczby.
func generator(nums ...int) <-chan int {
	outCh := make(chan int)
	go func(nums []int, ch chan<- int) {
		for _, n := range nums {
			ch <- n  // Wysyła liczbę do kanału
		}
		close(ch)  // Zamyka kanał po wysłaniu wszystkich liczb
	}(nums, outCh)
	return outCh
}

// squared: czyta liczby z inCh, podnosi je do kwadratu, a następnie wysyła wynik do innego kanału.
func squared(inCh <-chan int) <-chan int {
	outCh := make(chan int)
	go func(chIn <-chan int, chOut chan<- int) {
		for n := range chIn {
			chOut <- n * n  // Wysyła kwadrat liczby do kanału
		}
		close(chOut)  // Zamyka kanał po przetworzeniu wszystkich liczb
	}(inCh, outCh)
	return outCh
}

func main() {
	// Ustawienie pipeline
	nums := generator(1, 2, 3, 4)    // Generowanie liczb
	squares := squared(nums)         // Podnoszenie liczb do kwadratu

	// Odczyt i wydrukowanie kwadratów liczb
	for n := range squares {
		fmt.Println(n)
	}
}
