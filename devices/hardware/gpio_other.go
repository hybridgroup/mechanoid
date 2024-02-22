//go:build !tinygo

package hardware

func pinConfigure(pin Pin, config PinConfig) {
	// Not implemented
}

func pinSet(pin Pin, value bool) {
	// Not implemented
}

func pinGet(pin Pin) bool {
	// Not implemented
	return false
}
