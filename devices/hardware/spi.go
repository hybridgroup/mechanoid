package hardware

import (
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/orsinium-labs/wypes"
)

var _ engine.Device = &SPI{}

type SPIConfig struct{}

type SPI struct {
	Engine *engine.Engine
}

func NewSPIDevice(e *engine.Engine) *SPI {
	return &SPI{
		Engine: e,
	}
}

func (SPI) Init() error {
	return nil
}

func (SPI) Modules() wypes.Modules {
	// this is where the host machine's SPI would be initialized
	// and all the hosted functions setup

	return wypes.Modules{
		"machine": wypes.Module{
			"__tinygo_spi_configure": wypes.H4(SPIConfigure),
			"__tinygo_spi_transfer":  wypes.H2(SPITransfer),
		},
	}
}

func SPIConfigure(bus wypes.UInt8, sck wypes.UInt8, SDO wypes.UInt8, SDI wypes.UInt8) wypes.Void {
	// Not implemented
	return wypes.Void{}
}

func SPITransfer(bus wypes.UInt8, w wypes.UInt8) wypes.UInt8 {
	// Not implemented
	return 0
}
