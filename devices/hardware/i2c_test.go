package hardware

import (
	"testing"
)

func TestI2C(t *testing.T) {
	dev := I2C{}
	if dev.Init() != nil {
		t.Errorf("Init() returned non-nil")
	}

	if dev.Modules() == nil {
		t.Errorf("Modules() returned nil")
	}
}
