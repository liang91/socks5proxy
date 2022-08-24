package main

import "fmt"

type Person struct {
	name string
	age  int
}

var me Person

var Me = &me

func main() {
	if Me == nil {
		fmt.Println("Me is nil")
	} else {
		fmt.Println("Me is not nil", Me.name, Me.age)

	}
}
