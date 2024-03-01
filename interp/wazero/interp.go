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
	module  api.Module
	ctx     context.Context
}

func (i *Interpreter) Name() string {
	return "wazero"
}

func (i *Interpreter) Init() error {
	i.ctx = context.Background()
	conf := wazero.NewRuntimeConfigInterpreter()
	conf = conf.WithDebugInfoEnabled(false)
	conf = conf.WithMemoryLimitPages(1)
	i.runtime = wazero.NewRuntimeWithConfig(i.ctx, conf)
	return nil
}

func (i *Interpreter) DefineFunc(moduleName, funcName string, f any) error {
	panic("unsupported, use wzero.Modules instead")
}

func (i *Interpreter) Load(code []byte) error {
	var err error
	conf := wazero.NewModuleConfig()
	conf = conf.WithRandSource(cheapRand{})
	i.module, err = i.runtime.InstantiateWithConfig(i.ctx, code, conf)
	return err
}

func (i *Interpreter) Run() (engine.Instance, error) {
	var err error
	init := i.module.ExportedFunction("_initialize")
	if init != nil {
		_, err = init.Call(i.ctx)
		if err != nil {
			return nil, err
		}
	}
	return &Instance{i.module, i.ctx}, nil
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

// A fake RandSource for having fewer memory allocations.
//
// Should not be used with modules that do need an access to random functions.
type cheapRand struct{}

var _ io.Reader = cheapRand{}

func (cheapRand) Read(b []byte) (int, error) {
	return len(b), nil
}
