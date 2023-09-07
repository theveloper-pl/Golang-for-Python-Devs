package main

import (
	"fmt"
	"os"

)

func something() {
	fmt.Println("Defer odpala sie na koncu kodu")
}

func something2() {
	fmt.Println("W takiej kolejnosci")
}

func main() {


	file, err := os.Open("/home/ubuntu/tutorial/06-Defer_and_error_handling/example.txt")
	defer file.Close()
	defer something()
	defer something2()

	if err != nil {
		fmt.Println("Błąd podczas otwierania pliku:", err)
		return
	}



	// Czytanie z pliku, jakieś operacje...
	// ...

	fmt.Println("Operacje na pliku zostały zakończone")
	
}
