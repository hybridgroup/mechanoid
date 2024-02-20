//go:build tinygo

package main

import (
	"machine"

	"github.com/hybridgroup/tinywasm/filestore/flash"
)

var (
	console = machine.Serial
	fs      = &flash.FileStore{}
)

func dataStart() uint32 {
	return machine.FlashDataStart()
}

func dataEnd() uint32 {
	return machine.FlashDataEnd()
}
