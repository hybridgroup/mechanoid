package main

import (
	_ "embed"
	"time"

	"github.com/aykevl/board"
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/hybridgroup/mechanoid/interp/wasman"
)

//go:embed ping.wasm
var pingModule []byte

var (
	pingCount, pongCount int
)

func main() {
	display := NewDisplayDevice(board.Display.Configure())

	println("TinyWASM engine starting...")
	eng := engine.NewEngine()

	println("Using interpreter...")
	eng.UseInterpreter(&wasman.Interpreter{})

	println("Initializing engine...")
	eng.Init()

	if err := eng.Interpreter.DefineFunc("hosted", "pong", func() {
		pongCount++
		println("pong", pongCount)
		display.Pong(pongCount)
	}); err != nil {
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
		println("Ping", pingCount)
		ins.Call("ping")
		pingCount++
		display.Ping(pingCount)

		time.Sleep(1 * time.Second)
	}
}
