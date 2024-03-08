package hardware

import (
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/orsinium-labs/wypes"
)

var _ engine.Device = &GPIO{}

type Pin uint8
type PinMode uint8

const (
	PinInput PinMode = iota
	PinOutput
	PinInputPullup
	PinInputPulldown
)

type PinConfig struct {
	Mode PinMode
}

type GPIO struct{}

func (GPIO) Init() error {
	return nil
}

func (GPIO) Modules() wypes.Modules {
	// this is where the host machine's GPIO would be initialized
	// and all the hosted functions setup

	return wypes.Modules{
		"env": wypes.Module{
			"__tinygo_adc_read":       wypes.H1(ADCRead),
			"__tinygo_gpio_configure": wypes.H2(PinConfigure),
			"__tinygo_gpio_set":       wypes.H2(PinSet),
			"__tinygo_gpio_get":       wypes.H1(PinGet),
		},
	}
}

func PinConfigure(pin wypes.Int32, config wypes.Int32) wypes.Void {
	pinConfigure(Pin(pin), PinConfig{Mode: PinMode(config)})
	return wypes.Void{}
}

func PinSet(pin wypes.Int32, value wypes.Int32) wypes.Void {
	v := value != 0
	pinSet(Pin(pin), v)
	return wypes.Void{}
}

func PinGet(pin wypes.Int32) wypes.Int32 {
	if pinGet(Pin(pin)) {
		return 1
	}
	return 0
}
