package wasman

import (
	"github.com/hybridgroup/mechanoid"
	"github.com/hybridgroup/mechanoid/engine"

	wasmaneng "github.com/hybridgroup/wasman"
	"github.com/hybridgroup/wasman/config"
	"github.com/hybridgroup/wasman/wasm"
)

type Interpreter struct {
	linker   *wasmaneng.Linker
	module   *wasmaneng.Module
	instance *wasmaneng.Instance
	Memory   []byte

	references engine.ExternalReferences
}

func (i *Interpreter) Name() string {
	return "wasman"
}

func (i *Interpreter) Init() error {
	i.linker = wasmaneng.NewLinker(config.LinkerConfig{})

	// use host pre-allocated memory for instances
	if i.Memory != nil {
		if len(i.Memory)%65536 != 0 {
			return engine.ErrInvalidMemorySize
		}

		if err := i.linker.DefineMemory("env", "memory", i.Memory); err != nil {
			return err
		}
	}

	i.references = engine.NewReferences()

	return nil
}

func (i *Interpreter) Load(code []byte) error {
	conf := config.ModuleConfig{
		Recover: true,
		Logger:  mechanoid.Log,
	}

	var err error
	i.module, err = wasmaneng.NewModuleFromBytes(conf, code)
	if err != nil {
		return err
	}

	return nil
}

func (i *Interpreter) Run() (engine.Instance, error) {
	var err error
	i.instance, err = i.linker.Instantiate(i.module)
	if err != nil {
		return nil, err
	}

	_, _, err = i.instance.CallExportedFunc("_initialize")
	switch {
	case err == wasm.ErrExportedFuncNotFound:
		// no _initialize function, continue
	case err != nil:
		return nil, err
	}

	return &Instance{instance: i.instance}, nil
}

func (i *Interpreter) Halt() error {
	i.instance = nil
	return nil
}

// TODO: better implementation using generics?
func (i *Interpreter) DefineFunc(moduleName, funcName string, f any) error {
	switch tf := f.(type) {
	case func():
		return wasmaneng.DefineFunc(i.linker, moduleName, funcName, tf)
	case func() int32:
		return wasmaneng.DefineFunc01(i.linker, moduleName, funcName, tf)
	case func(int32):
		return wasmaneng.DefineFunc10(i.linker, moduleName, funcName, tf)
	case func(int32) int32:
		return wasmaneng.DefineFunc11(i.linker, moduleName, funcName, tf)
	case func(int32, int32):
		return wasmaneng.DefineFunc20(i.linker, moduleName, funcName, tf)
	case func(int32, int32) int32:
		return wasmaneng.DefineFunc21(i.linker, moduleName, funcName, tf)
	case func() uint32:
		return wasmaneng.DefineFunc01(i.linker, moduleName, funcName, tf)
	case func(uint32):
		return wasmaneng.DefineFunc10(i.linker, moduleName, funcName, tf)
	case func(uint32) uint32:
		return wasmaneng.DefineFunc11(i.linker, moduleName, funcName, tf)
	case func(uint32, uint32):
		return wasmaneng.DefineFunc20(i.linker, moduleName, funcName, tf)
	case func(uint32, uint32) uint32:
		return wasmaneng.DefineFunc21(i.linker, moduleName, funcName, tf)
	case func(uint32, uint32, uint32) uint32:
		return wasmaneng.DefineFunc31(i.linker, moduleName, funcName, tf)
	default:
		return engine.ErrInvalidFuncType
	}
}

func (i *Interpreter) MemoryData(ptr, sz uint32) ([]byte, error) {
	if i.instance.Memory == nil {
		return nil, engine.ErrMemoryNotDefined
	}
	if ptr+sz > uint32(len(i.instance.Memory.Value)) {
		return nil, engine.ErrMemoryOutOfBounds
	}

	return i.instance.Memory.Value[ptr : ptr+sz], nil
}

// References are the external references managed by the host module.
func (i *Interpreter) References() *engine.ExternalReferences {
	return &i.references
}
