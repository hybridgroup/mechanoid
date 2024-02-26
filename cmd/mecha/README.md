# `mecha` command line interface tool

## What it does

Command line interface to creating new Mechanoid projects, flashing boards, and more.

## How to install

You need to have TinyGo installed to use the `mecha` command.

See https://tinygo.org/getting-started/

In addition you need to install the `gonew` command:

```
go install golang.org/x/tools/cmd/gonew@latest
```

You can then install `mecha`:

```
go install github.com/hybridgroup/mechanoid/cmd/mecha@latest
```

## How to use

```
$ mecha
NAME:
   mecha - Mechanoid CLI

USAGE:
   mecha [global options] command [command options] 

VERSION:
   0.0.1

COMMANDS:
   new      create a new Mechanoid project
   build    build a Mechanoid project to a binary file
   flash    flash a Mechanoid project to a device
   test     run tests for a Mechanoid project
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## New project

```
mecha new project example.com/myproject
```


## New project based on template

```
mecha new project -t=blink example.com/myproject
```

## New module

```
mecha new module mymodule

```

## New module based on template

```
mecha new module -t=blink mymodule

```
