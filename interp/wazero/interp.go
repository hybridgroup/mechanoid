package wazero

import (
	"context"
	"errors"
	"io"

	"github.com/hybridgroup/mechanoid/engine"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

type Interpreter struct {
	runtime wazero.Runtime
	defs    map[string]map[string]any
	module  api.Module

	references engine.ExternalReferences
}

func (i *Interpreter) Name() string {
	return "wazero"
}

func (i *Interpreter) Init() error {
	ctx := context.Background()
	conf := wazero.NewRuntimeConfigInterpreter()
	conf = conf.WithDebugInfoEnabled(false)
	conf = conf.WithMemoryLimitPages(1)
	i.runtime = wazero.NewRuntimeWithConfig(ctx, conf)
	i.references = engine.NewReferences()
	return nil
}

func (i *Interpreter) DefineFunc(moduleName, funcName string, f any) error {
	if i.defs == nil {
		i.defs = make(map[string]map[string]any)
	}
	if _, exists := i.defs[moduleName]; !exists {
		i.defs[moduleName] = make(map[string]any)
	}
	i.defs[moduleName][funcName] = f
	return nil
}

func (i *Interpreter) Load(code engine.Reader) error {
	var err error
	ctx := context.Background()
	conf := wazero.NewModuleConfig()
	conf = conf.WithRandSource(cheapRand{})
	for moduleName, funcs := range i.defs {
		b := i.runtime.NewHostModuleBuilder(moduleName)
		for funcName, f := range funcs {
			b = b.NewFunctionBuilder().WithFunc(f).Export(funcName)
		}
		compiled, err := b.Compile(ctx)
		if err != nil {
			return err
		}
		_, err = i.runtime.InstantiateModule(ctx, compiled, conf)
		if err != nil {
			return err
		}
	}

	data, err := io.ReadAll(code)
	if err != nil {
		return err
	}
	i.module, err = i.runtime.InstantiateWithConfig(ctx, data, conf)
	return err
}

func (i *Interpreter) Run() (engine.Instance, error) {
	var err error
	ctx := context.Background()
	init := i.module.ExportedFunction("_initialize")
	if init != nil {
		_, err = init.Call(ctx)
		if err != nil {
			return nil, err
		}
	}
	return &Instance{i.module}, nil
}

func (i *Interpreter) Halt() error {
	ctx := context.Background()
	err := i.runtime.Close(ctx)
	i.runtime = nil
	return err
}

func (i *Interpreter) MemoryData(ptr, sz uint32) ([]byte, error) {
	memory := i.module.ExportedMemory("memory")
	if memory == nil {
		return nil, errors.New("memory not found")
	}
	data, inRange := memory.Read(ptr, sz)
	if !inRange {
		return nil, errors.New("out of range memory access")
	}
	return data, nil
}

// References are the external references managed by the host module.
func (i *Interpreter) References() *engine.ExternalReferences {
	return &i.references
}

// A fake RandSource for having fewer memory allocations.
//
// Should not be used with modules that do need an access to random functions.
type cheapRand struct{}

var _ io.Reader = cheapRand{}

func (cheapRand) Read(b []byte) (int, error) {
	return len(b), nil
}
