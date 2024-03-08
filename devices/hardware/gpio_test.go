package hardware

import (
	"testing"
)

func TestGPIO(t *testing.T) {
	dev := GPIO{}
	if dev.Init() != nil {
		t.Errorf("Init() returned non-nil")
	}

	if dev.Modules() == nil {
		t.Errorf("Modules() returned nil")
	}
}
