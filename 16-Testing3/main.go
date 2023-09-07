package main

import (
	"fmt"
	"time"
)



func main() {


	fmt.Println("Hello world")

}


func doSomethingVeryLong(a int, b int) int {
	time.Sleep(5 * time.Second)
	return a + b
}
