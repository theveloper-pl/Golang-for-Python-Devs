package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) Person { return Person{name, age} }
func (p Person) SayHello() {
	fmt.Println("Hello, my name is", p.name, "and I am", p.age, "years old!")
}

func randomAge() int { return rand.Intn(100) }
func main() {

	rand.Seed(time.Now().UnixNano())
	people := []Person{

		NewPerson("Alice", randomAge()),
		NewPerson("Bob", randomAge()),
		NewPerson("Charlie", randomAge()),
		NewPerson("Dave", randomAge())}
	for _, person := range people {
		if person.age > 50 {
			fmt.Println(person.name, "is older than 50")
		} else {
			fmt.Println(person.name, "is 50 or younger")
		}
		person.SayHello()
	}
}
