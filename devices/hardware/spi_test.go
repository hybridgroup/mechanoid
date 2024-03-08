package hardware

import (
	"testing"

	"github.com/hybridgroup/mechanoid/engine"
)

func TestSPI(t *testing.T) {
	eng := engine.NewEngine()
	dev := NewSPIDevice(eng)
	if dev == nil {
		t.Errorf("NewSPIDevice() returned nil")
	}

	if dev.Init() != nil {
		t.Errorf("Init() returned non-nil")
	}

	if dev.Modules() == nil {
		t.Errorf("Modules() returned nil")
	}
}
