//go:build !tinygo

package main

import "github.com/hybridgroup/tinywasm/filestore/filesystem"

var (
	console = UART{}
	fs      = &filesystem.FileStore{}
)

type UARTConfig struct{}

type UART struct {
}

// Configure the UART.
func (uart *UART) Configure(config UARTConfig) {
}

// Read from the UART.
func (uart *UART) Read(data []byte) (n int, err error) {
	return 0, nil
}

// Write to the UART.
func (uart *UART) Write(data []byte) (n int, err error) {
	return 0, nil
}

// Buffered returns the number of bytes currently stored in the RX buffer.
func (uart *UART) Buffered() int {
	return 0
}

// ReadByte reads a single byte from the UART.
func (uart *UART) ReadByte() (byte, error) {
	return 0, nil
}

// WriteByte writes a single byte to the UART.
func (uart *UART) WriteByte(b byte) error {
	return nil
}

func dataStart() uint32 {
	return 0
}

func dataEnd() uint32 {
	return 0
}
