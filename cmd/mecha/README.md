# `mecha` command line interface tool

## What it does

Command line interface to creating new Mechanoid projects, flashing boards, and more.

## How to install

You need to have TinyGo installed to use the `mecha` command.

See https://tinygo.org/getting-started/

In addition you need to install the `gonew` command:

```bash
go install golang.org/x/tools/cmd/gonew@latest
```

You can then install `mecha`:

```bash
go install github.com/hybridgroup/mechanoid/cmd/mecha@latest
```

## How to use

```bash
$ mecha
NAME:
   mecha - Mechanoid CLI

USAGE:
   mecha [global options] command [command options] 

VERSION:
   0.2.0-dev

COMMANDS:
   new      Create new Mechanoid project or module
   build    Build binary files for Mechanoid project and/or modules
   flash    Flash Mechanoid project to hardware
   run      Run code for Mechanoid project
   test     Run tests for Mechanoid project
   monitor  Monitor connection to hardware using the serial port
   about    About Mechanoid
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

### New project

```bash
mecha new project example.com/myproject
```

### New project based on template

```bash
mecha new project -t=blink example.com/myproject
```

### New module

```bash
mecha new module mymodule

```

### New TinyGo WASM module based on template

```bash
mecha new module -t=blink mymodule

```

### New Rust WASM module based on template

```bash
mecha new module -t=pingrs -type=rust pingrs

```

## New Zig WASM module based on template

```bash
mecha new module -t=pingzig -type=zig pingzig

```

### Build modules in current project

```bash
mecha build
```

or

```bash
mecha build modules
```

### Build application for current project

```bash
mecha build project
```

## Using `mecha` with Rust

If you want to use the `mecha` command with Rust, you will need to install Rust as follows.

- First install Rust by using the instructions here: https://www.rust-lang.org/tools/install

- Then, install the Rust `wasm32-unknown-unknown` target.

```bash
rustup target add wasm32-unknown-unknown
```
Any Rust modules in your project's `modules` directory should be automatically built when you run the `mecha build` command.

## Using `mecha` with Zig

If you want to use the `mecha` command with Zig, you will need to install Zig 0.11.0 by using the instructions here: https://ziglang.org/learn/getting-started/#installing-zig

Any Zig modules in your project's `modules` directory should be automatically built when you run the `mecha build` command.
