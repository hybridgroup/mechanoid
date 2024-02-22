# Architecture

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

## Application

The host application that the developer who uses Mechanoid is creating.

## Modules

The WASM modules that developers who are creating code for this Application are writing.

## Engine

The capabilities that the Application uses/exposes for Modules.

## Devices

Wrappers around specific devices such as displays or sensors that can be used by the Application and/or Modules.

## Network

Wrappers around specific networking capabilities such as WiFi or Bluetooth that can be used by the Application and/or Modules.

## Machine

Wrappers around low-level hardware interfaces such as GPIO or I2C that can be used by the Application and/or Modules.
