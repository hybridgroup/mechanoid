package main

import (
	"machine"

	_ "embed"
	"time"

	"github.com/hybridgroup/tinywasm/engine"
	"github.com/hybridgroup/tinywasm/filestore/flash"
	"github.com/hybridgroup/tinywasm/interp/wasman"
)

var (
	eng     *engine.Engine
	console = machine.Serial
)

func main() {
	time.Sleep(5 * time.Second)

	println("TinyWASM engine starting...")
	eng = engine.NewEngine()

	println("Using interpreter...")
	eng.UseInterpreter(&wasman.Interpreter{})

	println("Use file store...")
	eng.UseFileStore(&flash.FileStore{})

	println("Initializing engine...")
	eng.Init()

	println("Defining func...")
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
