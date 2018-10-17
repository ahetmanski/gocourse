package main

import (
	"fmt"
	"unsafe"
)

type myStruct struct {
	myFloat float64
	myBool  bool
	myInt   int32
}

func main() {
	fmt.Println(unsafe.Sizeof(&myStruct{}))

	s := myStruct{}
	fmt.Println(unsafe.Offsetof(s.myBool))
	fmt.Println(unsafe.Offsetof(s.myInt))
	fmt.Println(unsafe.Sizeof(myStruct{}))
}
