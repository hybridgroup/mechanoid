package hardware

import (
	"testing"
)

func TestADC(t *testing.T) {
	dev := ADC{}
	if dev.Init() != nil {
		t.Errorf("Init() returned non-nil")
	}

	if dev.Modules() == nil {
		t.Errorf("Modules() returned nil")
	}
}
