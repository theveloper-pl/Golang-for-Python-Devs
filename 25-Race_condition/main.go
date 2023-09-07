package main

import (
	"fmt"
	"sync"
)
/*
"Race condition", czyli "warunek wyścigu", występuje, gdy dwa lub więcej wątków (lub gorutyn w Go) 
próbuje jednocześnie uzyskać dostęp i modyfikować współdzielony zasób (np. zmienną lub strukturę danych) 
w sposób, który zależy od tego, w jaki sposób przeplatają się operacje na tym zasobie. 
Innymi słowy, jest to sytuacja, w której wynik operacji zależy od nieprzewidywalnej sekwencji zdarzeń.
*/

var balance int = 0
var wg sync.WaitGroup
var mutex sync.Mutex  // Definicja zmiennej mutex

func deposit(amount int) {
    mutex.Lock()  // Blokada mutex przed dostępem do zasobu

    currentBalance := balance   // odczytaj bieżący bilans
    currentBalance += amount    // dodaj wartość do bieżącego bilansu
    balance = currentBalance    // aktualizuj globalny bilans

    mutex.Unlock()  // Odblokowanie mutex po zakończeniu pracy z zasobem

    wg.Done()  // informacja o zakończeniu pracy przez gorutynę
}

func main() {
    clients := 1000
    amount := 100

    // Ustawienie licznika WaitGroup na liczbę klientów
    wg.Add(clients)

    for i := 0; i < clients; i++ {
        go deposit(amount)  // uruchomienie funkcji wpłaty dla każdego klienta jako osobnej gorutyny
    }

    // Oczekiwanie na zakończenie wszystkich gorutyn
    wg.Wait()

    fmt.Printf("Stan końcowy konta: %d\n", balance)
}



// go run -race nazwa_pliku.go