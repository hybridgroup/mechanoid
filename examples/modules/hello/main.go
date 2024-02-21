//go:build tinygo

package main

//go:export hola
func hola(msg string) uint32

const msg = "Hello, WebAssembly!"

var msgdata [64]byte

//go:export hello
func hello() {
	copy(msgdata[:], []byte(msg))
	hola(string(msgdata[:len(msg)]))
}

func main() {}
