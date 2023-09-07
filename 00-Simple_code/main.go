package main

import (
	"fmt"
)


func main() {
	var name string
	fmt.Println("Jak masz na imie?")
	fmt.Scanln(&name)
	welcome := fmt.Sprintf("Hello, %s! \n", name)
	fmt.Println(welcome)
}


func init() {
	fmt.Println("Init startuje przed main()")
}

func init() {
	fmt.Println("Mozemy miec wiele funkcji init()")
}

func init() {
	fmt.Println("Bedą one wywołane w kolejności w jakiej zostały zadeklarowane")
}

/*

Wnioski:
- init() jest wywoływane przed main(). Nie musisz ręcznie wywoływać tych funkcji. 
Obie funkcje sa wywolywane automatycznie. Funkcji init() nie musisz nawet deklrarować w kodzie.
- Każda zmienna podczas kompilacji ma określony typ nawet jeśli nie jest on podany jawnie
- Moduł fmt należy do standardowej biblioteki w Go, zawiera podstawowe funkcje wejścia/wyjścia
- Każdy plik .go musi zawierać ,,package X,, . Dodatkowo package main musi mieć funkcję main() - init() jest opcjonalne
- W golangu powszechne są ,,pointery,, / ,,wskazniki,,
	Funkcje Scan, Scanf i Scanln wymagają wskaźników do zmiennych jako argumentów, ponieważ 
	wprowadzają zmiany bezpośrednio w pamięci tych zmiennych

*/
