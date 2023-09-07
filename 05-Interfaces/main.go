package main

import (
	"fmt"
)

/*

Go używa interfejsów do definiowania zachowań, które typ musi implementować. 
Nie wymaga jawnego deklarowania, że dany typ implementuje dany interfejs. 
Jeśli typ ma metody wymagane przez interfejs, automatycznie go implementuje.

*/

// Interfejs definiujący metodę Speak dla różnych zwierząt.
type Animal interface {
	Speak() string
	Move()
}


// Funkcja, która przyjmuje interfejs Animal.
func IntroduceAnimal(a Animal) {
	fmt.Println(a.Speak())
}


// Typ dla psa.
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Hau, hau!"
}


// Typ dla kota.
type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Miau!"
}



func main() {
	dog := Dog{Name: "Burek"}
	cat := Cat{Name: "Mruczek"}

	IntroduceAnimal(dog)  // Wydrukuje: Hau, hau!
	IntroduceAnimal(cat)  // Wydrukuje: Miau!
}
