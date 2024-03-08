package hardware

import (
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/orsinium-labs/wypes"
)

var _ engine.Device = &I2C{}

type I2CConfig struct{}

type I2C struct{}

func (I2C) Init() error {
	return nil
}

func (i2c I2C) Modules() wypes.Modules {
	// this is where the host machine's I2C would be initialized
	// and all the hosted functions setup

	return wypes.Modules{
		"machine": wypes.Module{
			"__tinygo_i2c_configure":     wypes.H3(I2CConfigure),
			"__tinygo_i2c_set_baud_rate": wypes.H2(I2CSetBaudRate),
			"__tinygo_i2c_transfer":      wypes.H5(I2CTransfer),
		},
	}
}

func I2CConfigure(bus wypes.UInt8, scl wypes.UInt8, sda wypes.UInt8) wypes.Void {
	// Not implemented
	return wypes.Void{}
}

func I2CSetBaudRate(bus wypes.UInt8, br wypes.UInt32) wypes.Void {
	// Not implemented
	return wypes.Void{}
}

func I2CTransfer(bus wypes.UInt8, w wypes.Byte, wlen wypes.Int, r wypes.Byte, rlen wypes.Int) wypes.Int {
	// Not implemented
	return 0
}
