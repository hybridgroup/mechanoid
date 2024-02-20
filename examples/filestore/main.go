package main

import (
	_ "embed"
	"time"

	"github.com/hybridgroup/tinywasm/engine"
	"github.com/hybridgroup/tinywasm/interp/wasman"
)

var (
	eng *engine.Engine
)

func main() {
	time.Sleep(5 * time.Second)

	println("TinyWASM engine starting...")
	eng = engine.NewEngine()

	println("Using interpreter...")
	eng.UseInterpreter(&wasman.Interpreter{})

	println("Use file store...")
	eng.UseFileStore(fs)

	println("Initializing engine...")
	eng.Init()

	if err := eng.Interpreter.DefineFunc("hosted", "pong", pongFunc); err != nil {
		println(err.Error())
		return
	}

	// start up CLI
	cli()
}

func pongFunc() {
	println("pong")
}
