//go:build tinygo

package hardware

import "machine"

func pinConfigure(pin Pin, config PinConfig) {
	mode := machine.PinOutput
	switch config.Mode {
	case PinOutput:
		mode = machine.PinOutput
	case PinInput:
		mode = machine.PinInput
	case PinInputPullup:
		mode = machine.PinInputPullup
	case PinInputPulldown:
		mode = machine.PinInputPulldown
	}
	machine.Pin(pin).Configure(machine.PinConfig{machine.PinMode(mode)})
}

func pinSet(pin Pin, value bool) {
	machine.Pin(pin).Set(value)
}

func pinGet(pin Pin) bool {
	return machine.Pin(pin).Get()
}
