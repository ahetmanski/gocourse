package main

import "fmt"

func main() {
	var a int = 8
	b := 3
	var c float32

	c = float32(a) / float32(b)
	fmt.Printf("Value of c: %f\n", c)
}
