# TinyWASM

TinyWASM is a WebAssembly runtime environment for embedded microcontrollers written using TinyGo.


## How to use

### Simple

Loads an embedded WASM module and then runs it by calling the exported `Ping()` function:

```go
package main

import (
	_ "embed"
	"time"

	"github.com/hybridgroup/tinywasm/engine"
	"github.com/hybridgroup/tinywasm/interp/wasman"
)

//go:embed ping.wasm
var pingModule []byte

func main() {
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
```


## Architecture

```mermaid
flowchart TD
    subgraph Application
        App
    end
    App-->Engine
    subgraph Modules
        WASM1
        WASM2
    end
    subgraph Engine
        FileStore
        Interpreter
        Devices
    end
    FileStore-->Modules
    Interpreter-->Modules
    Interpreter-->Devices
    Devices--->Machine
    Devices--->Hardware
    Devices--->Network
    subgraph Hardware
        Sensor
        Displayer
        LEDSetter
    end
    subgraph Network
        Net
        Bluetooth
    end
    subgraph Machine
        GPIO
        ADC
        I2C
        SPI
    end
    Displayer-->SPI
    Sensor-->GPIO
    Sensor-->I2C
```

#### Application

The host application that the developer who uses TinyWASM is creating.

#### Modules

The WASM modules that developers who are creating code for this Application are writing.

#### Engine

The capabilities that the Application uses/exposes for Modules.

#### Devices

Wrappers around specific devices such as displays or sensors that can be used by the Application and/or Modules.

#### Network

Wrappers around specific networking capabilities such as WiFi or Bluetooth that can be used by the Application and/or Modules.

#### Machine

Wrappers around low-level hardware interfaces such as GPIO or I2C that can be used by the Application and/or Modules.

## Goals

- [X] Able to run small WASM modules designed for specific embedded runtime interfaces.
- [X] Hot loading/unloading of WASM modules.
- [X] Local storage system for WASM modules.
- [ ] Allow the engine to be used/extended for different embedded application use cases, e.g. CLI, WASM4 runtime, others.
- [ ] Configurable system to allow the bridge interface to host capabilities to be defined per application.
