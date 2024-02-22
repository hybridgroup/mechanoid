![Mechanoid logo](https://mechanoid.io/images/logo-black.png)

Mechanoid is a framework for developing applications using WebAssembly for embedded microcontrollers written using TinyGo.

## Simple

### WebAssembly guest program

```go
//go:build tinygo

package main

//go:wasmimport hosted pong
func pong()

//go:export ping
func ping() {
	pong()
}

func main() {}
```

Compile this program to WASM using TinyGo:

```
$ tinygo build -size short -o ./modules/ping/ping.wasm -target ./modules/ping/wasm-unknown.json -no-debug ./modules/ping
   code    data     bss |   flash     ram
      9       0       0 |       9       0
```

### Mechanoid host application

The Mechanoid host application that runs on the hardware, loads this WASM module and then runs it by calling the exported `Ping()` function:

```go
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
	println("Mechanoid engine starting...")
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
```

You can compile and flash the WASM runtime engine and the WASM program onto an Adafruit PyBadge (an ARM 32-bit microcontroller with 192k of RAM) with this command:

```
$ tinygo flash -size short -target pybadge -monitor ./examples/simple
   code    data     bss |   flash     ram
 101012    1736   72216 |  102748   73952
Connected to /dev/ttyACM0. Press Ctrl-C to exit.
Mechanoid engine starting...
Using interpreter...
Initializing engine...
Initializing interpreter...
Initializing devices...
Defining func...
Loading module...
Running module...
building index space
initializing memory
initializing functions
initializing globals
running start func
Calling ping...
pong
Calling ping...
pong
Calling ping...
pong
Calling ping...
pong
```

More examples are available here:
https://github.com/hybridgroup/mechanoid-examples

## Getting started

- Install the Mechanoid command line tool
- Create a new project
- Make something amazing!

## `mecha` command line tool

```
NAME:
   mecha - Mechanoid WASM embedded development tools

USAGE:
   mecha [global options] command [command options] 

COMMANDS:
   new      create a new Mechanoid project
   flash    flash a Mechanoid project to a device
   test     run tests for a Mechanoid project
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

## How it works

See [ARCHITECTURE.md](./ARCHITECTURE.md) for more information.

## Goals

- [X] Able to run small WASM modules designed for specific embedded runtime interfaces.
- [X] Hot loading/unloading of WASM modules.
- [X] Local storage system for WASM modules.
- [ ] Allow the engine to be used/extended for different embedded application use cases, e.g. CLI, WASM4 runtime, others.
- [ ] Configurable system to allow the bridge interface to host capabilities to be defined per application.
