//go:build !(wasman || epsilon)

package interp

import (
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/hybridgroup/mechanoid/interp/epsilon"
)

func NewInterpreter() engine.Interpreter {
	return &epsilon.Interpreter{}
}
