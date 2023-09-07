package main

import (
	"errors"
	"fmt"
	"os"
	"io/ioutil"
	"strconv"
	"net/http"
)

// Divide dzieli dwie liczby i zwraca wynik.
// Jeśli dzielnik wynosi 0, funkcja zwraca błąd.
func Divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return dividend / divisor, nil
}

func main() {
	// Poprawne dzielenie:
	result, err := Divide(4, 2)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("4 / 2 =", result)

	// Dzielenie przez zero:
	result, err = Divide(4, 0)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("4 / 0 =", result)
}




func do_something() {
	// Otwieranie pliku
	file, err := os.Open("nieistniejący_plik.txt")
	if err != nil {
		fmt.Println("Błąd podczas otwierania pliku:", err)
		return
	}
	defer file.Close()

	// Dalej kod używający pliku...
}


func do_something2() {
	// Zapis do pliku
	data := []byte("Witaj, świecie!")
	err := ioutil.WriteFile("przykład.txt", data, 0644)
	if err != nil {
		fmt.Println("Błąd podczas zapisu do pliku:", err)
		return
	}
}


func do_something3() {
	// Parsowanie liczby całkowitej z ciągu znaków:
	str := "123a"
	val, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Błąd podczas parsowania '%s' jako int: %v\n", str, err)
		return
	}
	fmt.Println("Wartość:", val)
}



func do_something4() {
	// Wysłanie żądania HTTP
	resp, err := http.Get("https://www.example.com")
	if err != nil {
		fmt.Println("Błąd podczas wysyłania żądania:", err)
		return
	}
	defer resp.Body.Close()

	// Dalej kod przetwarzający odpowiedź...
}

