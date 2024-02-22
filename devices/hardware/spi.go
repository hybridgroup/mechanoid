package hardware

import (
	"github.com/hybridgroup/mechanoid/engine"
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

func (s *SPI) Init() error {
	// this is where the host machine's SPI would be initialized
	// and all the hosted functions setup
	if s.Engine == nil {
		return engine.ErrInvalidEngine
	}

	if err := s.Engine.Interpreter.DefineFunc("machine", "__tinygo_spi_configure", SPIConfigure); err != nil {
		println(err.Error())
		return err
	}

	if err := s.Engine.Interpreter.DefineFunc("machine", "__tinygo_spi_transfer", SPITransfer); err != nil {
		println(err.Error())
		return err
	}

	return nil
}

func SPIConfigure(bus uint8, sck Pin, SDO Pin, SDI Pin) {
	// Not implemented
}

func SPITransfer(bus uint8, w uint8) uint8 {
	// Not implemented
	return 0
}
