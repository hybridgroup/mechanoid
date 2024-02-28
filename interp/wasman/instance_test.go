package wasman

import (
	"testing"

	"github.com/hybridgroup/mechanoid/interp/tester"
)

func TestInstance(t *testing.T) {
	tester.InstanceTest(t, &Interpreter{
		Memory: make([]byte, 65536),
	})
}
