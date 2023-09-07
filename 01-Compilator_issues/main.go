package main

import (
	"testing" // Każdy zaimportowany modul musi zostać wykorzystany - inaczej kompilator zwróci błąd
)

func main() {
	name := "Adas"  // Każda utworzona zmienna musi zostać wykorzystana - inaczej kompilator zwróci błąd




// W pakiecie "mypackage"
type myType struct {
    field int
}
// W głównym pakiecie
var v mypackage.myType
fmt.Println(v.field)  // Błąd: "field" jest nieosiągalne  - To jak "_" oraz "__" w Pythonie


}


