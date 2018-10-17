package main

import "fmt"

func main() {
	foo := [3]int{1, 2, 3}
	fmt.Println(foo, len(foo))

	bar := [...]int{2, 3, 5, 7}
	fmt.Println(bar, len(bar))
}
