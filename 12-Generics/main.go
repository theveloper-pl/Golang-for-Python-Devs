package main

import (
	"fmt"
)





/*
 Comparable types are:
 - Numeric types (int, float, etc.)
 - String
 - Bool
 - Pointers
 - Structs containing only comparable types

Not comparable are:
 - slices
 - maps
 - functions
 - channels
 - Structs containing not comparable types
*/





type Person struct {
	name string
	age  int
}

func print_something_comparable[T comparable](x T) {
	fmt.Println(x)
} 

func print_something_any[T any](x T) {
	fmt.Println(x)
}


/*

 Comparable oraz any są ,,wbudowanymi,, interfejsami. Na podobnej zasadzie możemy tworzyć własne ,,ograniczniki,,
 
*/

func xyz[T interface {}](x T) {
	fmt.Println(x)
}


func main() {

	p := Person{name: "Jan", age:  42,}

	fmt.Println("Using comparable:")
	print_something_comparable(1) 
	print_something_comparable("a")  
//	print_something_comparable([]int{1, 2, 3})  // Error
	print_something_comparable(p) 

	fmt.Println("Using any:")
	print_something_any(1) 
	print_something_any("a")
	print_something_any([]int{1, 2, 3})
	print_something_any(p)
}
