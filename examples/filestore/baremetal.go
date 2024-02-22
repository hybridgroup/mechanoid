//go:build tinygo

package main

import (
	"machine"

	"github.com/hybridgroup/mechanoid/filestore/flash"
)

var (
	console = machine.Serial
	fs      = &flash.FileStore{}
)

func dataStart() uint32 {
	return uint32(machine.FlashDataStart())
}

func dataEnd() uint32 {
	return uint32(machine.FlashDataEnd())
}
