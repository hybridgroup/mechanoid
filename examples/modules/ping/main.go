//go:build tinygo

package main

//go:wasmimport hosted pong
func pong()

//go:export ping
func ping() {
	pong()
}

func main() {}
