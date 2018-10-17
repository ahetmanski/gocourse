package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	bob := person{"Bob", 20}
	alice := &person{name: "Alice"}

	fmt.Println(alice.name, alice.age)

	bob.age = 32
	fmt.Println(bob)

	alice.age = 51
	fmt.Println(alice)
}
