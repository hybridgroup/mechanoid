package hardware

import (
	"testing"
)

func TestSPI(t *testing.T) {
	dev := SPI{}
	if dev.Init() != nil {
		t.Errorf("Init() returned non-nil")
	}

	if dev.Modules() == nil {
		t.Errorf("Modules() returned nil")
	}
}
