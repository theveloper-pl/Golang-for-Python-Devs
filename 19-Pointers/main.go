package main

import (
	"fmt"
)

type User struct {
	email string
	username string
	age int
}

func (u User) Email() string{
	return u.email
}



// func (u *User) SetEmail(email string){
// 	u.email = email
// }

func (u User) SetEmail(email string){
	u.email = email
}

func main() {

	user := User{
		email: "email@gmail.com",
		username: "username",
		age: 20,
	}

	fmt.Println(user.Email())
	user.SetEmail("changedemail")
	fmt.Println(user.Email())

}


func returnUserByValue() User {
	user := User{
		email: "email@gmail.com",
		username: "username",
		age: 20,
	}
	return user
}


func returnUserByPointer() *User {
	user := User{
		email: "email@gmail.com",
		username: "username",
		age: 20,
	}
	return &user
}


//  |
//  |
//  |
//  |
//  |
// \ /
//  V



























// Stack VS Heap

// By value are stored in stack
// By pointer are stored in heap

