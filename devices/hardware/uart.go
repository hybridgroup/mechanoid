package hardware

import (
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/orsinium-labs/wypes"
)

var _ engine.Device = &UART{}

type UARTConfig struct{}

type UART struct {
	eng *engine.Engine
}

func NewUARTDevice(e *engine.Engine) *UART {
	return &UART{
		eng: e,
	}
}

func (*UART) Init() error {
	return nil
}

func (*UART) Modules() wypes.Modules {
	// this is where the host machine's UART would be initialized
	// and all the hosted functions setup
	return wypes.Modules{
		"machine": wypes.Module{
			"__tinygo_uart_configure": wypes.H3(UARTConfigure),
			"__tinygo_uart_read":      wypes.H3(UARTRead),
			"__tinygo_uart_write":     wypes.H3(UARTWrite),
		},
	}
}

func UARTConfigure(bus wypes.UInt8, tx wypes.UInt8, rx wypes.UInt8) wypes.Void {
	// Not implemented
	return wypes.Void{}
}

func UARTRead(bus wypes.UInt8, buf wypes.Byte, bufLen wypes.UInt32) wypes.UInt32 {
	// Not implemented
	return 0

}

func UARTWrite(bus wypes.UInt8, buf wypes.Byte, bufLen wypes.UInt32) wypes.UInt32 {
	// Not implemented
	return 0
}
