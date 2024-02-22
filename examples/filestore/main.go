package main

import (
	_ "embed"
	"time"

	"github.com/hybridgroup/mechanoid/engine"
	"github.com/hybridgroup/mechanoid/interp/wasman"
)

var (
	eng *engine.Engine
)

func main() {
	time.Sleep(5 * time.Second)

	println("TinyWASM engine starting...")
	eng = engine.NewEngine()

	println("Using interpreter...")
	eng.UseInterpreter(&wasman.Interpreter{
		Memory: make([]byte, 65536),
	})

	println("Use file store...")
	eng.UseFileStore(fs)

	println("Initializing engine...")
	eng.Init()

	if err := eng.Interpreter.DefineFunc("hosted", "pong", pongFunc); err != nil {
		println(err.Error())
		return
	}

	if err := eng.Interpreter.DefineFunc("env", "hola", holaFunc); err != nil {
		println(err.Error())
		return
	}

	// start up CLI
	cli()
}

func pongFunc() {
	println("pong")
}

func holaFunc(ptr uint32, size uint32) uint32 {
	msg, err := eng.Interpreter.MemoryData(ptr, size)
	if err != nil {
		println(err.Error())
		return 0
	}
	println(string(msg))
	return size
}
