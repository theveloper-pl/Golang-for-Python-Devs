package main

import (
	"fmt"
)

/*

Multiple return values

*/


func main() {
	value, err := do_something("Hello world")
	fmt.Println(value, err)

	value2 := do_something_different()
	fmt.Println(value2)
}


func do_something(something string) (int, error) {
	fmt.Println(something)
    return 123456, nil
}


func do_something_different() (value int) {
	value = 2137
    return
}