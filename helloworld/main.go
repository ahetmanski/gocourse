package main

import (
	"runtime"

	"github.com/ahetmanski/gocourse/helloworld/hello"
)

func main() {
	hello.Hello(runtime.GOARCH)
}
