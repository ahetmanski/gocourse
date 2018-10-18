package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m, "len:", len(m))

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	delete(m, "k2")
	fmt.Println("map:", m)

	_, ok := m["k2"]
	fmt.Println("k2 in map:", ok)

	n := map[string]int{"Level": 1, "NextLevel": 2}
	fmt.Println("map:", n)
}
