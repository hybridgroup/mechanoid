package hardware

import (
	"github.com/hybridgroup/mechanoid/engine"
)

var _ engine.Device = &UART{}

type UARTConfig struct{}

type UART struct {
	Engine *engine.Engine
}

func NewUARTDevice(e *engine.Engine) *UART {
	return &UART{
		Engine: e,
	}
}

func (uart *UART) Init() error {
	// this is where the host machine's UART would be initialized
	// and all the hosted functions setup
	if uart.Engine == nil {
		return engine.ErrInvalidEngine
	}

	if err := uart.Engine.Interpreter.DefineFunc("machine", "__tinygo_uart_configure", UARTConfigure); err != nil {
		println(err.Error())
		return err
	}

	if err := uart.Engine.Interpreter.DefineFunc("machine", "__tinygo_uart_read", UARTRead); err != nil {
		println(err.Error())
		return err
	}

	if err := uart.Engine.Interpreter.DefineFunc("machine", "__tinygo_uart_write", UARTWrite); err != nil {
		println(err.Error())
		return err
	}

	return nil
}

func UARTConfigure(bus uint8, tx Pin, rx Pin) {
	// Not implemented
}

func UARTRead(bus uint8, buf *byte, bufLen int) int {
	// Not implemented
	return 0

}

func UARTWrite(bus uint8, buf *byte, bufLen int) int {
	// Not implemented
	return 0
}
