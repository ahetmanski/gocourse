package main

import "fmt"

func main() {
	s1 := []byte{1, 2, 3, 5}
	fmt.Printf("%p, %v, %d, %d\n", s1, s1, len(s1), cap(s1))

	s2 := []byte{7, 11, 13}
	s1 = append(s1, s2...)
	fmt.Printf("%p, %v, %d, %d\n", s1, s1, len(s1), cap(s1))

	s1 = append(s1, 23)
	fmt.Printf("%p, %v, %d, %d\n", s1, s1, len(s1), cap(s1))

	s1 = append(s1, 29)
	fmt.Printf("%p, %v, %d, %d\n", s1, s1, len(s1), cap(s1))

	s3 := make([]int, 10, 100)
	fmt.Printf("Slice#3 %p, %v, len=%d, cap=%d\n", s3, s3, len(s3), cap(s3))

	const (
		from, to = 2, 5
	)
	s1 = append(s1[:from], s1[to:]...)
	fmt.Printf("Slice#1 %p, %v, %d, %d\n", s1, s1, len(s1), cap(s1))

	s4 := s1[from:to]
	fmt.Printf("Before changes %p s1=%v\n", s1, s1)
	s4[0] = 253
	s4[1] = 254
	s4[2] = 255
	fmt.Printf("After changes %p s1=%v, %p s4=%v\n", s1, s1, s4, s4)
}
