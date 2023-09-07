package main

import (
	"testing"
	"fmt"
)


//  go test -bench=.


func BenchmarkReturnUserByValue(b *testing.B) {
	user := User{}
	for i:=0; i<b.N; i++ {
		user = returnUserByValue()
	}
	fmt.Println(user)
}

func BenchmarkReturnUserByPointer(b *testing.B) {
	user := &User{}
	for i:=0; i<b.N; i++ {
		user = returnUserByPointer()
	}
	fmt.Println(user)
}













