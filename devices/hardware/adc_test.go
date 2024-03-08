package hardware

import (
	"testing"

	"github.com/hybridgroup/mechanoid/engine"
)

func TestADC(t *testing.T) {
	eng := engine.NewEngine()
	dev := NewADCDevice(eng)
	if dev == nil {
		t.Errorf("NewADCDevice() returned nil")
	}

	if dev.Init() != nil {
		t.Errorf("Init() returned non-nil")
	}

	if dev.Modules() == nil {
		t.Errorf("Modules() returned nil")
	}
}
