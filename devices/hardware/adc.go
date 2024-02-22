package hardware

import (
	"github.com/hybridgroup/mechanoid/engine"
)

var _ engine.Device = &ADC{}

type ADC struct {
	Engine *engine.Engine
}

func (a *ADC) Init() error {
	// this is where the host machine's GPIO would be initialized
	// and all the hosted functions setup
	if a.Engine == nil {
		return engine.ErrInvalidEngine
	}

	if err := a.Engine.Interpreter.DefineFunc("machine", "__tinygo_adc_read", ADCRead); err != nil {
		println(err.Error())
		return err
	}

	return nil
}

func ADCRead(pin Pin) uint16 {
	// Not implemented
	return 0
}
