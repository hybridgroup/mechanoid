//go:build wazero

package interp

import (
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/hybridgroup/mechanoid/interp/wazero"
)

func NewInterpreter() engine.Interpreter {
	return &wazero.Interpreter{}
}
