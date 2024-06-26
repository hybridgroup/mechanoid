package wazero

import (
	"context"
	"errors"
	"fmt"
	"io"
	"runtime"

	"github.com/hybridgroup/mechanoid"
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/orsinium-labs/wypes"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

type Interpreter struct {
	runtime wazero.Runtime
	module  api.Module
	modules wypes.Modules
}

func (i *Interpreter) Name() string {
	return "wazero"
}

func (i *Interpreter) Init() error {
	return i.init()
}

func (i *Interpreter) init() error {
	mechanoid.DebugMemory("Interpreter Init")

	ctx := context.Background()
	conf := wazero.NewRuntimeConfigInterpreter()
	conf = conf.WithDebugInfoEnabled(false)
	conf = conf.WithMemoryLimitPages(1)
	i.runtime = wazero.NewRuntimeWithConfig(ctx, conf)
	return nil
}

func (i *Interpreter) SetModules(modules wypes.Modules) error {
	mechanoid.Log("Registering host modules...")

	if i.modules == nil {
		i.modules = modules
		return nil
	}
	for modName, funcs := range modules {
		_, found := i.modules[modName]
		if !found {
			i.modules[modName] = funcs
			continue
		}
		for funcName, funcDef := range funcs {
			i.modules[modName][funcName] = funcDef
		}
	}
	return nil
}

func (i *Interpreter) Load(code engine.Reader) error {
	if i.runtime == nil {
		if err := i.init(); err != nil {
			return fmt.Errorf("init wazero runtime: %v", err)
		}
	}

	mechanoid.DebugMemory("Interpreter Load")

	err := i.defineModules()
	if err != nil {
		return fmt.Errorf("register wazero host modules: %v", err)
	}
	ctx := context.Background()
	conf := wazero.NewModuleConfig()
	conf = conf.WithRandSource(cheapRand{})
	data, err := io.ReadAll(code)
	if err != nil {
		return fmt.Errorf("read wasm binary: %v", err)
	}
	i.module, err = i.runtime.InstantiateWithConfig(ctx, data, conf)
	return err
}

func (i *Interpreter) defineModules() error {
	if i.modules == nil {
		return nil
	}
	refs := wypes.NewMapRefs()
	for modName, mod := range i.modules {
		err := i.defineModule(modName, mod, refs)
		if err != nil {
			return fmt.Errorf("define module %s: %v", modName, err)
		}
	}
	return nil
}

func (i *Interpreter) defineModule(modName string, m wypes.Module, refs wypes.Refs) error {
	mb := i.runtime.NewHostModuleBuilder(modName)
	for funcName, funcDef := range m {
		fb := mb.NewFunctionBuilder()
		fb = fb.WithGoModuleFunction(
			wazeroAdaptHostFunc(funcDef, refs),
			funcDef.ParamValueTypes(),
			funcDef.ResultValueTypes(),
		)
		mb = fb.Export(funcName)
	}
	ctx := context.Background()
	compiled, err := mb.Compile(ctx)
	if err != nil {
		return err
	}
	conf := wazero.NewModuleConfig()
	conf = conf.WithRandSource(cheapRand{})
	_, err = i.runtime.InstantiateModule(ctx, compiled, conf)
	return err
}

func wazeroAdaptHostFunc(hf wypes.HostFunc, refs wypes.Refs) api.GoModuleFunction {
	return api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
		adaptedStack := wypes.SliceStack(stack)
		store := wypes.Store{
			Memory:  mod.Memory(),
			Stack:   &adaptedStack,
			Refs:    refs,
			Context: ctx,
		}
		hf.Call(store)
	})
}

func (i *Interpreter) Run() (engine.Instance, error) {
	mechanoid.DebugMemory("Interpreter Run")

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
	mechanoid.DebugMemory("Interpreter Halt")

	ctx := context.Background()
	err := i.runtime.Close(ctx)
	i.runtime = nil
	i.module = nil

	// force a garbage collection to free memory
	runtime.GC()
	mechanoid.DebugMemory("Interpreter Halt after GC")

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
