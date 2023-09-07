package main

import (
	"fmt"
)



func main() {


	fmt.Println("Hello world")

}


//go test -fuzz=FuzzDoSomething -fuzztime 5s -v
func doSomething(a int, b int) int {
	fmt.Println(a, b)
	return a + b
}
