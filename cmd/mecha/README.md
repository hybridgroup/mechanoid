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

Use `export GOPRIVATE` (only needed until the repo is public)

```bash
export GOPRIVATE=github.com/hybridgroup/mecha*
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
   0.0.1-dev

COMMANDS:
   new      Create new Mechanoid project or module
   build    Build binary files for Mechanoid project
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

## New project

```bash
mecha new project example.com/myproject
```


## New project based on template

```bash
mecha new project -t=blink example.com/myproject
```

## New module

```bash
mecha new module mymodule

```

## New module based on template

```bash
mecha new module -t=blink mymodule

```
