package wazero

import (
	"testing"

	"github.com/hybridgroup/mechanoid/interp/tester"
)

func TestInstance(t *testing.T) {
	tester.InstanceTest(t, &Interpreter{})
}
