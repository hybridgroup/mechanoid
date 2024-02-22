package hardware

import (
	"github.com/hybridgroup/mechanoid/engine"
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

type GPIO struct {
	Engine *engine.Engine
}

const moduleName = "env"

func NewGPIODevice(e *engine.Engine) *GPIO {
	return &GPIO{
		Engine: e,
	}
}

func (g *GPIO) Init() error {
	// this is where the host machine's GPIO would be initialized
	// and all the hosted functions setup
	if g.Engine == nil {
		return engine.ErrInvalidEngine
	}

	if err := g.Engine.Interpreter.DefineFunc(moduleName, "__tinygo_gpio_configure", PinConfigure); err != nil {
		println(err.Error())
		return err
	}

	if err := g.Engine.Interpreter.DefineFunc(moduleName, "__tinygo_gpio_set", PinSet); err != nil {
		println(err.Error())
		return err
	}

	if err := g.Engine.Interpreter.DefineFunc(moduleName, "__tinygo_gpio_get", PinGet); err != nil {
		println(err.Error())
		return err
	}

	return nil
}

func PinConfigure(pin int32, config int32) {
	pinConfigure(Pin(pin), PinConfig{Mode: PinMode(config)})
}

func PinSet(pin int32, value int32) {
	v := value != 0
	pinSet(Pin(pin), v)
}

func PinGet(pin int32) int32 {
	if pinGet(Pin(pin)) {
		return 1
	}
	return 0
}
