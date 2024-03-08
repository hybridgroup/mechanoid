package hardware

import (
	"testing"
)

func TestUART(t *testing.T) {
	dev := UART{}
	if dev.Init() != nil {
		t.Errorf("Init() returned non-nil")
	}

	if dev.Modules() == nil {
		t.Errorf("Modules() returned nil")
	}
}
