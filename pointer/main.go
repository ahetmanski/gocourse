package main

import "fmt"

func main() {
	var sampleVariable int = 140
	var Pointer *int

	if Pointer == nil {
		fmt.Printf("Uninitialized Pointer stores next address: %v\n", Pointer)
	} else {
		fmt.Printf("nil *Pointer\n", *Pointer)
	}

	Pointer = &sampleVariable

	fmt.Printf("Address of sampleVariable value: %p\n", &sampleVariable)
	fmt.Printf("Address stored in Pointer variable: %p\n", Pointer)
	fmt.Printf("Value of *Pointer variable: %d\n", *Pointer)
}
