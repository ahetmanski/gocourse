package main

import (
	"fmt"
	"unsafe"
)

type myStruct struct {
	myBool  bool
	myFloat float64
	myInt   int32
}

func main() {
	fmt.Println(unsafe.Sizeof(&myStruct{}))

	s := myStruct{}
	fmt.Println(unsafe.Offsetof(s.myBool))
	fmt.Println(unsafe.Offsetof(s.myInt))
	fmt.Println(unsafe.Sizeof(myStruct{}))
}
