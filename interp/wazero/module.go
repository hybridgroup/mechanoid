package wazero

import (
	"context"

	"github.com/stealthrocket/wazergo"
	"github.com/tetratelabs/wazero"
)

// AddModule registers a set of host-defined functions in the interpreter.
func AddModule[State wazergo.Module](
	interp *Interpreter,
	name string,
	state State,
	funcs wazergo.Functions[State],
) error {
	mod := hostModule[State]{
		name:  name,
		funcs: funcs,
		state: state,
	}

	ctx := context.Background()
	builder := wazergo.Build[State](interp.runtime, mod)
	compiled, err := builder.Compile(ctx)
	if err != nil {
		return err
	}
	conf := wazero.NewModuleConfig()
	conf = conf.WithRandSource(cheapRand{})
	_, err = interp.runtime.InstantiateModule(ctx, compiled, conf)

	inst := wazergo.MustInstantiate(ctx, interp.runtime, mod)
	interp.ctx = wazergo.WithModuleInstance[State](interp.ctx, inst)
	return err
}

type hostModule[State wazergo.Module] struct {
	name  string
	funcs wazergo.Functions[State]
	state State
}

func (m hostModule[State]) Name() string {
	return m.name
}

func (m hostModule[State]) Functions() wazergo.Functions[State] {
	return m.funcs
}

func (m hostModule[State]) Instantiate(
	ctx context.Context,
	options ...wazergo.Option[State],
) (State, error) {
	return m.state, nil
}
