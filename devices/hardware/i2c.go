package hardware

import (
	"github.com/hybridgroup/mechanoid/engine"
)

var _ engine.Device = &I2C{}

type I2CConfig struct{}

type I2C struct {
	Engine *engine.Engine
}

func NewI2CDevice(e *engine.Engine) *I2C {
	return &I2C{
		Engine: e,
	}
}

func (i2c *I2C) Init() error {
	// this is where the host machine's I2C would be initialized
	// and all the hosted functions setup
	if i2c.Engine == nil {
		return engine.ErrInvalidEngine
	}

	if err := i2c.Engine.Interpreter.DefineFunc("machine", "__tinygo_i2c_configure", I2CConfigure); err != nil {
		println(err.Error())
		return err
	}

	if err := i2c.Engine.Interpreter.DefineFunc("machine", "__tinygo_i2c_set_baud_rate", I2CSetBaudRate); err != nil {
		println(err.Error())
		return err
	}

	if err := i2c.Engine.Interpreter.DefineFunc("machine", "__tinygo_i2c_transfer", I2CTransfer); err != nil {
		println(err.Error())
		return err
	}

	return nil
}

func I2CConfigure(bus uint8, scl Pin, sda Pin) {
	// Not implemented
}

func I2CSetBaudRate(bus uint8, br uint32) {
	// Not implemented
}

func I2CTransfer(bus uint8, w *byte, wlen int, r *byte, rlen int) int {
	// Not implemented
	return 0
}
