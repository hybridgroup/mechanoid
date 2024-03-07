//go:build !(wasman || wazero)

package interp

import (
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/hybridgroup/mechanoid/interp/wasman"
)

func NewInterpreter() engine.Interpreter {
	return &wasman.Interpreter{}
}
