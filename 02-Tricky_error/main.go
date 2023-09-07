package main

import (
	"fmt"
	"time"
)

func main() {
	s := []int{1, 2, 3}
	for _, v := range s {          // w Golangu jest tylko pętla for
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(2 * time.Second)

}