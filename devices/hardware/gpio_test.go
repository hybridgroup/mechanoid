package hardware

import (
	"testing"

	"github.com/hybridgroup/mechanoid/engine"
)

func TestGPIO(t *testing.T) {
	eng := engine.NewEngine()
	dev := NewGPIODevice(eng)
	if dev == nil {
		t.Errorf("NewGPIODevice() returned nil")
	}

	if dev.Init() != nil {
		t.Errorf("Init() returned non-nil")
	}

	if dev.Modules() == nil {
		t.Errorf("Modules() returned nil")
	}
}
