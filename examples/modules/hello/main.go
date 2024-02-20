package main

import "unsafe"

var (
	x int
)

//go:export hello
func hello(ptr uint32, size uint32) uint32 {
	x++
	msg := "Hello, World " + string(x)

	if len(msg) > int(size) {
		msg = msg[:size]
	}
	copy(*(*[]byte)(unsafe.Pointer(&ptr)), msg)

	return uint32(len(msg))
}

func main() {}
