package main

import (
	_ "embed"
	"time"

	"github.com/hybridgroup/tinywasm/engine"
)

//go:embed ping.wasm
var binaryModule []byte

func main() {
	eng := engine.NewEngine()
	eng.Init()

	if err := eng.Interpreter.DefineFunc("hosted", "pong", pongFunc); err != nil {
		println(err.Error())
		return
	}

	mod, err := eng.Interpreter.Load(binaryModule)
	if err != nil {
		println(err.Error())
		return
	}

	if err := eng.Interpreter.Run(); err != nil {
		println(err.Error())
		return
	}

	for {
		mod.Call("ping")

		time.Sleep(1 * time.Second)
	}
}

func pongFunc() {
	println("pong")
}
