package main

import (
	_ "embed"
	"time"

	"github.com/hybridgroup/mechanoid/engine"
	"github.com/hybridgroup/mechanoid/interp/wasman"
)

//go:embed ping.wasm
var pingModule []byte

func main() {
	time.Sleep(5 * time.Second)

	println("TinyWASM engine starting...")
	eng := engine.NewEngine()

	println("Using interpreter...")
	eng.UseInterpreter(&wasman.Interpreter{})

	println("Initializing engine...")
	eng.Init()

	println("Defining func...")
	if err := eng.Interpreter.DefineFunc("hosted", "pong", pongFunc); err != nil {
		println(err.Error())
		return
	}

	println("Loading module...")
	if err := eng.Interpreter.Load(pingModule); err != nil {
		println(err.Error())
		return
	}

	println("Running module...")
	ins, err := eng.Interpreter.Run()
	if err != nil {
		println(err.Error())
		return
	}

	for {
		println("Calling ping...")
		ins.Call("ping")

		time.Sleep(1 * time.Second)
	}
}

func pongFunc() {
	println("pong")
}
