package hardware

import (
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/orsinium-labs/wypes"
)

var _ engine.Device = &ADC{}

type ADC struct {
	eng *engine.Engine
}

func NewADCDevice(e *engine.Engine) *ADC {
	return &ADC{
		eng: e,
	}
}

func (*ADC) Init() error {
	return nil
}

func (*ADC) Modules() wypes.Modules {
	return wypes.Modules{
		"machine": wypes.Module{
			"__tinygo_adc_read": wypes.H1(ADCRead),
		},
	}
}

func ADCRead(pin wypes.UInt8) wypes.UInt16 {
	// Not implemented
	return 0
}
