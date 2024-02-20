package wasman

import (
	"github.com/hybridgroup/tinywasm/engine"

	wasmaneng "github.com/c0mm4nd/wasman"
	"github.com/c0mm4nd/wasman/config"
	"github.com/c0mm4nd/wasman/wasm"
)

var heapMemory = make([]byte, 65536)

type Interpreter struct {
	linker   *wasmaneng.Linker
	module   *wasmaneng.Module
	instance *wasmaneng.Instance
}

func (i *Interpreter) Name() string {
	return "wasman"
}

func (i *Interpreter) Init() error {
	i.linker = wasmaneng.NewLinker(config.LinkerConfig{})

	if err := i.linker.DefineMemory("env", "memory", heapMemory); err != nil {
		return err
	}

	return nil
}

func (i *Interpreter) Load(code []byte) error {
	conf := config.ModuleConfig{
		Recover: true,
		Logger:  i.Log,
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
	return nil
}

// TODO: better implementation using generics?
func (i *Interpreter) DefineFunc(moduleName, funcName string, f interface{}) error {
	switch f.(type) {
	case func():
		if err := wasmaneng.DefineFunc(i.linker, moduleName, funcName, f.(func())); err != nil {
			return err
		}
		return nil
	case func() int32:
		if err := wasmaneng.DefineFunc01(i.linker, moduleName, funcName, f.(func() int32)); err != nil {
			return err
		}
		return nil
	case func(int32):
		if err := wasmaneng.DefineFunc10(i.linker, moduleName, funcName, f.(func(int32))); err != nil {
			return err
		}
		return nil
	case func(int32) int32:
		if err := wasmaneng.DefineFunc11(i.linker, moduleName, funcName, f.(func(int32) int32)); err != nil {
			return err
		}
		return nil
	case func(int32, int32) int32:
		if err := wasmaneng.DefineFunc21(i.linker, moduleName, funcName, f.(func(uint32, uint32) uint32)); err != nil {
			return err
		}
		return nil
	case func() uint32:
		if err := wasmaneng.DefineFunc01(i.linker, moduleName, funcName, f.(func() uint32)); err != nil {
			return err
		}
		return nil
	case func(uint32):
		if err := wasmaneng.DefineFunc10(i.linker, moduleName, funcName, f.(func(uint32))); err != nil {
			return err
		}
		return nil
	case func(uint32) uint32:
		if err := wasmaneng.DefineFunc11(i.linker, moduleName, funcName, f.(func(uint32) uint32)); err != nil {
			return err
		}
		return nil
	case func(uint32, uint32) uint32:
		if err := wasmaneng.DefineFunc21(i.linker, moduleName, funcName, f.(func(uint32, uint32) uint32)); err != nil {
			return err
		}
		return nil
	default:
		return engine.ErrInvalidFuncType
	}
}

func (i *Interpreter) Log(msg string) {
	println(msg)
}
