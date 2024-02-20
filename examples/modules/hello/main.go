package main

import (
	"unsafe"

	"github.com/hybridgroup/tinywasm/convert"
)

var (
	x int
)

//go:export hello
func hello(ptr uint32, size uint32) uint32 {
	x++
	msg := "Hello, World " + convert.IntToString(x)

	if len(msg) > int(size) {
		msg = msg[:size]
	}
	copy(*(*[]byte)(unsafe.Pointer(&ptr)), msg)

	return uint32(len(msg))
}

func main() {}
