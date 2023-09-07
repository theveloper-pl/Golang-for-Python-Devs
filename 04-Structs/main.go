package main

import (
	"encoding/json"
	"fmt"
)


/*


Typowanie w Go jest statyczne. 
Oznacza to, że typ każdej zmiennej musi być jasno zdefiniowany podczas kompilacji i nie może zostać 
zmieniony w czasie działania programu. 
Struktury w Go mają ściśle określone pola o określonych typach, które są sprawdzane podczas kompilacji. 
Nie możemy dynamicznie dodawać lub usuwać pól z istniejącej struktury w trakcie działania programu.
Dzięki statycznemu typowaniu błędy związane z niezgodnością typów są wykrywane na etapie kompilacji. 
Daje to pewność co do tego, jakie pola i metody posiada dany obiekt oraz jakiego są typu.

*/


type Adres struct {
    Ulica, Miasto string
}
type Osoba struct {
    Imię   string
    Adres  // Anonimowe osadzenie - W Golangu nie mamy klasycznego dziedziczenia.
	//  Zamiast tego wykorzystuje się kompozycje oraz interfejsy
}


// Tagi struktur to metadane dodawane do definicji pól struktury w Go. - 
// Bardzo często wykorzystywane przez zewnętrzne biblioteki np GORM lub GIN
type Person struct {
	FirstName string `json:"first_name"`    // W Golangu o enkapsulacji decyduje wielkość pierwszej litery -
											// 	jezeli jest duza to pole jest publiczne, jezeli mała to prywatne dla pakietu
	LastName  string `json:"last_name"`
	Age       int    `json:"age,omitempty"`
	Nickname  string `json:"-"`
}


// Funkcja działająca jak konsktruktor
func NowaOsoba(FirstName string, LastName string, Nickname string, Age int) *Person {
    return &Person{FirstName: FirstName, LastName: LastName, Nickname: Nickname, Age: Age}
}

// Struktury nie mają wbudowanych metod jak klasy w Pythonie, ale możemy twórzyć funkcję z tzw. 
// receiverem która będzie działać jak metoda
func (p Person) PelneDane() string {
	return fmt.Sprintf("Imię: %s, Nazwisko: %s, Wiek: %d", p.FirstName ,p.LastName, p.Age)
}

func main() {
	p := Person{
		FirstName: "Jan",
		LastName:  "Kowalski",
		Nickname:  "JK",
		// Uwaga: Age jest pominięte, więc będzie miało wartość zerową (0)
	}

	// Serializacja do JSON
	jsonData, err := json.Marshal(p) // Multiple return values
	if err != nil {
		fmt.Println("Błąd podczas serializacji:", err)
		return
	}

	// Wydrukowanie sformatowanego JSON
	fmt.Println(string(jsonData))

	// Pokazanie, że pole Nickname jest pominięte mimo że ma wartość
	fmt.Println("Nickname:", p.Nickname)

	fmt.Print(p.PelneDane())
}

