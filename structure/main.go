package main

import "fmt"

type person struct {
	name string
	m    metrics
}

type metrics struct {
	age int
}

func main() {
	bob := person{"Bob", metrics{20}}
	alice := &person{name: "Alice"}

	fmt.Println(alice.name, alice.m.age)

	bob.m.age = 32
	fmt.Println(bob)

	alice.m.age = 51
	fmt.Println(alice)

	(*alice).m.age = 62
	fmt.Println(alice)
}
