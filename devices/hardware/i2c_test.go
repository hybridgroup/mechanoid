package hardware

import (
	"testing"

	"github.com/hybridgroup/mechanoid/engine"
)

func TestI2C(t *testing.T) {
	eng := engine.NewEngine()
	dev := NewI2CDevice(eng)
	if dev == nil {
		t.Errorf("NewI2CDevice() returned nil")
	}

	if dev.Init() != nil {
		t.Errorf("Init() returned non-nil")
	}

	if dev.Modules() == nil {
		t.Errorf("Modules() returned nil")
	}
}
