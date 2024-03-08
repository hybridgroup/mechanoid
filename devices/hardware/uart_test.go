package hardware

import (
	"testing"

	"github.com/hybridgroup/mechanoid/engine"
)

func TestUART(t *testing.T) {
	eng := engine.NewEngine()
	dev := NewUARTDevice(eng)
	if dev == nil {
		t.Errorf("NewUARTDevice() returned nil")
	}

	if dev.Init() != nil {
		t.Errorf("Init() returned non-nil")
	}

	if dev.Modules() == nil {
		t.Errorf("Modules() returned nil")
	}
}
