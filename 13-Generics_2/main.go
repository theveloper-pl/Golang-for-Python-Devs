package main

import (
	"fmt"
)

// Użycie ~ przed każdym z tych typów oznacza, że Number może być reprezentowany przez dowolny z tych konkretnych typów, zarówno całkowitych, jak i zmiennoprzecinkowych.
// Bez tego do funkcji  Min mozna by korzystac tylko i wylaczni z typu Number co niema sensu bo Number jest interfejsem

// Number to interfejs typu "constraint", który ogranicza możliwe typy wartości
// do liczbowych typów zarówno całkowitych, jak i zmiennoprzecinkowych.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

// Point to nowy typ danych oparty na typie int.
type Point int

// Min jest funkcją generyczną, która przyjmuje dwa argumenty dowolnego typu 
// liczbowego (dzięki interfejsowi Number) i zwraca mniejszą z nich.
func Min[T Number](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func main() {
	// Tworzenie dwóch zmiennych typu Point
	x, y := Point(5), Point(2)

	// Wywołanie funkcji Min z dwoma argumentami typu Point
	// i wydrukowanie wyniku.
	fmt.Println(Min(x, y))
}
